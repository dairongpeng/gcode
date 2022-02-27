# 访问echo-go的svc

## kubernetes网络结构

> Master和Node网络 -> Service网络 -> Pod网络 -> 容器网络

### Kubernetes 所在节点的网络
这级的网络是最容易理解的，就是 Master 和 Node 节点的 IP 地址

查看节点信息：
```shell
$ kubectl get nodes

NAME          STATUS   ROLES    AGE    VERSION
kube-master   Ready    master   131d   v1.15.0
kube-node-1   Ready    <none>   131d   v1.15.0
kube-node-2   Ready    <none>   131d   v1.15.0
```

找到所有的节点信息之后,找到每个节点的 IP 地址
```shell
$ kubectl get nodes kube-master -o 'jsonpath={.status.addresses[0].address}'

10.192.0.2

$ kubectl get nodes kube-node-1 -o 'jsonpath={.status.addresses[0].address}'

10.192.0.3

$ kubectl get nodes kube-node-2 -o 'jsonpath={.status.addresses[0].address}'

10.192.0.4
```

**Tips: 可以把 Kubernetes 的 Master 和 Node 以 Docker 容器的方式启动，可以实现在单台物理机上面创建一个 Kubernetes 集群的功能**

### Service网络
其中 Service 的网络也称为 Cluster 的网络。
就是整个 Kubernetes 创建的集群叫做 Cluster，在这个 Cluster 内部会有很多的 Service，
这些 Service 通过网络转发把流量转发给自己负责的 Pod，然后进入 Pod 内部的容器中。

由于Pod是在不停变变化的，Ip不固定，这时抽象出来一个Service,然后给 Service 分配一个固定的虚拟 IP，这个虚拟 IP 就叫做 ClusterIP。
应用程序通过 ClusterIP 访问 Service，然后通过 Service 的转发最终访问到 Pod 里面的服务

```shell
$ kubectl get svc -n devops -o wide

NAME          TYPE       CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE     SELECTOR
echo-go-svc   NodePort   10.109.20.42   <none>        9090:31147/TCP   5m27s   app=echo-go
```

在创建 Service 的时候，我们注意到在 spec 节点里面定义了一个 type: NodePort 的键值对，这个键值对为我们访问 Service 提供了可能。
这个 NodePort 表示在创建 Service 的时候，将在 Kubernetes 的 Nodes 节点上绑定一个端口映射，映射到 Service 的端口。
这个也就是我们上面看到的查看 Service 的输出的时候，为什么会有一列 PORTS。我们看一下那个 PROTS 的列：
```shell
$ kubectl get svc -n devops -o wide

NAME          TYPE       CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE     SELECTOR
echo-go-svc   NodePort   10.109.20.42   <none>        9090:31147/TCP   6m54s   app=echo-go
```

既然在 Kubernetes 的 Nodes 节点上面已经有端口映射了，那么可以访问我们刚得到的 Nodes 的 IP 试试看。
在上面的 PORTS 列中，左边的端口是 Service 的端口9090，右边是 Node 上绑定的映射端口31147：（这里的端口号需要根据上面的信息填写）

```shell
$ curl http://10.192.0.3:31147
GET /

User-Agent: curl/7.47.0
Accept: */*


$ curl http://10.192.0.4:31147
GET /

User-Agent: curl/7.47.0
Accept: */*
```

Service是如何维护其背后代理的一组pod的呢，其实这种关系的维护是通过 Endpoints 资源来进行的。
当每个 Service 创建的时候，都会创建一个同名的 Endpoints 资源，在这个资源中维护了 Service 的端口到后端 Pod IP 和端口的映射关系。
当后端的 Pod 因为销毁重建导致 Pod IP 发生变化的情况下，Endpoints 资源会同步更新到新的 Pod IP。
这样 Service 就可以及时感知到最新的变化，从而不影响请求的转发功能。

可以使用如下的命令来查看下这个同名的 Endpoints 资源：
```shell
$ kubectl get endpoints -n devops

NAME          ENDPOINTS         AGE
echo-go-svc   10.244.2.3:9090   21m
```

可以看到确实存在一个 Endpoints 资源和我们刚刚创建的 Service 名称一样。
我们可以继续探究下 Endpoints 内部是如何维护 Service 和 Pod IP 之间的关系的：
```shell
kubectl get endpoints echo-go-svc -n devops -o yaml
```

```yaml
apiVersion: v1
kind: Endpoints
metadata:
  annotations:
    endpoints.kubernetes.io/last-change-trigger-time: "2020-12-28T08:53:53Z"
  creationTimestamp: "2020-12-28T08:53:53Z"
  labels:
    app: echo-go
  name: echo-go-svc
  namespace: devops
  resourceVersion: "2394"
  selfLink: /api/v1/namespaces/devops/endpoints/echo-go-svc
  uid: 34d4dac7-3286-41dc-a172-36f74e96d477
subsets:
- addresses:
  - ip: 10.244.2.3
    nodeName: kube-node-1
    targetRef:
      kind: Pod
      name: echo-go-pod1
      namespace: devops
      resourceVersion: "2011"
      uid: c91b647e-1dce-4b2d-987e-707f9137ca39
  ports:
  - name: tcp-9090-9090-echo-go-http
    port: 9090
    protocol: TCP
```

从这个 Endpoints 资源的定义中发现，存在一个 subsets 的节点，在这个节点里面通过 Node 的信息和 targetRef 节点定义了到具体 Pod 的映射关系。
为了证明我们上面说过的这个 Endpoints 资源会根据后端 Pod IP 动态变化的情况而进行更新，我们可以再创建一个 Pod，让 Service 可以转发请求到两个 Pod。


