# 给echo-go容器创建service

## 编写service的yaml文件
```yaml
apiVersion: v1
kind: Service
metadata:
  name: echo-go-svc
  labels:
    app: echo-go
spec:
  ports:
    - name: tcp-9090-9090-echo-go-http
      port: 9090
      protocol: TCP
      targetPort: 9090
  selector:
    app: echo-go
  type: NodePort
```

## 解释
- port 表示 Service 对外提供的端口
- targetPort 表示 Pod 暴露的端口
- selector根据pod中定义的labels来选中当前service要服务的pid

## 使用yaml文件创建service
```shell
$ kubectl create -f echo-go-svc.yaml -n devops

service/echo-go-svc created
```

## 查看创建的service
```shell
$ kubectl get svc -n devops -o wide

NAME          TYPE       CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE   SELECTOR
echo-go-svc   NodePort   10.109.20.42   <none>        9090:31147/TCP   27s   app=echo-go
```