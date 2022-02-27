# 使用容器镜像echo制作pod

## 创建自己的命名空间
```shell
$ kubectl create ns devops

namespace/devops created
```

## 使用yaml文件创建pod
```shell
$ kubectl create \-f echo\-go\-pod1.yaml \-n devops

pod/echo\-go\-pod1 created
```

## 查看创建的pod
```shell
kubectl get pods \-n devops

NAME READY STATUS RESTARTS AGE

echo\-go\-pod1 1/1 Running 0 7s
```

## 查看pod创建的过程信息
```shell
kubectl describe pods \-n devops
```

