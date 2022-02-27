package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const (
	DockerClientVersion = "1.37"
)

// 为了突出 Docker 镜像构建的部分功能，我们这里的 Tar 文件使用手动打好的指定的 Tar 包。
// Tar包的内容是应用的二进制文件和Dockerfile两个文件压缩的。
// tar cvf echo-go.tar Dockerfile echo-go // 压缩两个文件
//$ tar tf echo-go.tar // 解压缩tar
//echo-go
//Dockerfile
func main() {
	var tarFile string
	var imageName string
	flag.StringVar(&tarFile, "tar-file", "", "tar file to build docker image")
	flag.StringVar(&imageName, "image-name", "", "dest image name")
	flag.Parse()
	if tarFile == "" || imageName == "" {
		fmt.Println("Err: no tar file or dest image name specified")
		return
	}

	// 创建镜像构建的 Client 对象
	imageBuildClient, err := client.NewClientWithOpts(client.WithVersion(DockerClientVersion))
	if err != nil {
		fmt.Println("Err: create docker build client error,", err.Error())
		return
	}

	// 打开 Tar 文件
	tarFileFp, err := os.Open(tarFile)
	if err != nil {
		fmt.Println("Err: open tar file error,", err.Error())
		return
	}

	defer tarFileFp.Close()

	// 发送构建请求
	ctx := context.Background()

	imageBuildResp, err := imageBuildClient.ImageBuild(ctx, tarFileFp, types.ImageBuildOptions{
		Tags:       []string{imageName},
		Dockerfile: "./Dockerfile",
	})
	if err != nil {
		fmt.Println("Err: send image build request error,", err.Error())
		return
	}
	defer imageBuildResp.Body.Close()

	// 打印构建输出
	_, err = io.Copy(os.Stdout, imageBuildResp.Body)
	if err != nil {
		fmt.Println("Err: read image build response error,", err.Error())
		return
	}
}
