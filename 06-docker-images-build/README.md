# image-build使用方法

## 编译image-build.go文件
```shell
go build image-build.go
```

## 运行image-build 且指定目标tar文件
- 执行
```shell
./image-build -tar-file 'echo-go.tar' -image-name 'registry.cn-hangzhou.aliyuncs.com/chenshi-kubernetes/echo-go:1.0'
```
- 日志
```shell
{"stream":"Step 1/4 : FROM alpine:latest"}
{"stream":"\n"}
{"stream":" ---\u003e b7b28af77ffe\n"}
{"stream":"Step 2/4 : RUN mkdir -p /home/app/"}
{"stream":"\n"}
{"stream":" ---\u003e Running in 53ef77286821\n"}
{"stream":" ---\u003e 762c131b4b92\n"}
{"stream":"Step 3/4 : ADD ./echo-go /home/app/"}
{"stream":"\n"}
{"stream":" ---\u003e c1944d8f132a\n"}
{"stream":"Step 4/4 : CMD /home/app/echo-go -host 0.0.0.0 -port 9090"}
{"stream":"\n"}
{"stream":" ---\u003e Running in 530c50a64728\n"}
{"stream":" ---\u003e e526a3ff2dc1\n"}
{"aux":{"ID":"sha256:e526a3ff2dc102cac7ca1a092f393b5a4a5555c005ae10ab0836f36f255d0f39"}}
{"stream":"Successfully built e526a3ff2dc1\n"}
{"stream":"Successfully tagged registry.cn-hangzhou.aliyuncs.com/chenshi-kubernetes/echo-go:1.0\n"}
```

- 完成后查看构建的镜像
```shell
$ docker images registry.cn-hangzhou.aliyuncs.com/chenshi-kubernetes/echo-go
REPOSITORY                                                     TAG                 IMAGE ID            CREATED             SIZE
registry.cn-hangzhou.aliyuncs.com/chenshi-kubernetes/echo-go   1.0                 4e1f086cdba9
```

## 运行测试的镜像
```shell
docker run -p 9090:9090 registry.cn-hangzhou.aliyuncs.com/chenshi-kubernetes/echo-go:1.0
```

## 测试运行的容器
```shell
$ curl http://localhost:9090/it/ping -H 'Content-Type: text/plain'
GET /it/ping

User-Agent: curl/7.47.0
Accept: */*
Content-Type: text/plain
```