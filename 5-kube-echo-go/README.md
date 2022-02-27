# echo-go编译

## 跨平台编译方法
```shell
# mac上编译linux和windows二进制
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build server.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build server.go

# linux上编译mac和windows二进制
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build server.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build server.go

# windows上编译mac和linux二进制
SET CGO_ENABLED=0 SET GOOS=darwin SET GOARCH=amd64 go build server.go
SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build server.go
```

- 这里我们在macos上编译linux上使用的二进制
```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o echo-go .
```

## 把编译好的echo-go和Dockerfile共同打成tar，供镜像构建使用
```shell
tar cvf echo-go.tar Dockerfile echo-go
```

## 压缩好tar文件后，进入docker-build中查看具体构建脚本
```shell
cd ../docker-build
```