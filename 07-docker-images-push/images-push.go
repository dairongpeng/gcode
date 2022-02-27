package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
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

// 需要特别注意的是，要用正确的可访问镜像中心的用户名和密码替换上面代码中的 xxx。
// 到目前为止，我们已经为在 Kubernetes 中发布第一个容器应用做好了全部的准备。接下来将进入 Kubernetes 相关的知识应用阶段。
func main() {
	var imageName string
	flag.StringVar(&imageName, "image-name", "", "image name to push")
	flag.Parse()
	if imageName == "" {
		fmt.Println("Err: no image name specified")
		return
	}

	// 创建镜像推送的 Client 对象
	imagePushClient, err := client.NewClientWithOpts(client.WithVersion(DockerClientVersion))
	if err != nil {
		fmt.Println("Err: create docker push client error,", err.Error())
		return
	}

	// 构建镜像推送的鉴权信息
	imagePushAuthConfig := types.AuthConfig{
		Username: "xxx",
		Password: "xxx",
	}
	imagePushAuth, _ := json.Marshal(&imagePushAuthConfig)

	// 发送镜像推送的请求
	ctx := context.Background()
	imagePushResp, err := imagePushClient.ImagePush(ctx, imageName, types.ImagePushOptions{
		RegistryAuth: base64.URLEncoding.EncodeToString(imagePushAuth),
	})
	if err != nil {
		fmt.Println("Err: send image push request error,", err.Error())
		return
	}
	defer imagePushResp.Close()

	// 打印镜像推送的输出
	_, err = io.Copy(os.Stdout, imagePushResp)
	if err != nil {
		fmt.Println("Err: read image push response error,", err.Error())
		return
	}
}
