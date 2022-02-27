package cmds

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	k8s "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
)

// 以下这些值在每次实验中因为环境重建问题可能发现变化，需要自行填写
// 获取这些值的方法参考第三节我们讲解创建 ServiceAccount 名称为
// shiyanlou-admin的地方。
var (
	// K8SCertificateData 表示 Kubernetes 服务端证书
	K8SCertificateData = `-----BEGIN CERTIFICATE-----
MIICyDCCAbCgAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl
cm5ldGVzMB4XDTE4MDUyMTA4NTgwOFoXDTI4MDUxODA4NTgwOFowFTETMBEGA1UE
AxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKi6
iiiSNO8TK3eYFsaCJF3ZlZreno3z/4cFHjY7C5Bct1ZaEUB15xO8MlBt3puEpb5o
o2fZ2+IIL1NRWyvjAsexJ4oyHk0xBP0a9KjEseypiw+m5lxe826GKLUX18BguaPX
Dge8qIh7b3/zWEfYkb7G/tLjtNKIDYDN+OOt6tjohjZ7FSP8G1qXEuj7MaqjFq4a
LB3uSkzITZ4aOPP0Yrpa9dzSjq+hHWz6H88Tg98oZL7PRIrrHOPXhMYXZwfoEtGM
dr73Abze5+2tLdN5Nv+5mtXbLxpVL+x6mBxwGEQ2bqUspuJ/SgHBrbV5ylmPzffp
vk1WfNy8k6ZDf2I+DU0CAwEAAaMjMCEwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB
/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBABq1HL188anTft34C9HG9i7rMqmO
Uyzv1YiFeQNaO0jp7IlbzN8RO6PV0tT/SJbCh3MmwP/SatrrpNTI4snwm9aTIqRv
bSmOo63uu2yPQPosUvPtsNF8XmzXVF3vqMcdy/J/w9hcwRZVJ53K/M8noc3rwZ9d
O2k2wXa4F7Q3bIIgcQhXgHiWT2iGi1n61Rnci+PAePZNkX1X1DhNPz/6UVToxxur
i0v7L8KEvTjprdNYML/aZCHVwamvz9y6dFJ3SIGok34EGSWWd6SoJPD3dd7Now8V
oByo9DPJ5u6y45iAb7kj7NEc15P1Qq94srB6Zs5B4qzAW2uzJL0YhThpRWA=
-----END CERTIFICATE-----`

	// K8SAPIServer 表示 Kubernetes API Server 地址
	K8SAPIServer = "https://10.192.0.2:6443"

	// K8SAPIToken 表示 ServiceAccount shiyanlou-admin 的 Secret 对应的 Token
	K8SAPIToken = `eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6InNoaXlhbmxvdS1hZG1pbi10b2tlbi1mcWo5dyIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJzaGl5YW5sb3UtYWRtaW4iLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiI3Mjk4NmRjMy0xZTM1LTExZTktODU4My1jYTQwZWVlNDBmOTgiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6ZGVmYXVsdDpzaGl5YW5sb3UtYWRtaW4ifQ.T6SKaRlHo3G_DItrRxNbKS6UHbe3-HmxjmSkgUd6JQWhWll5Y0VUX8AFCHjATzPFSjmKNmE9va_ljemm4H06O8a0AF4NTUZoliodhghquomtuS0nVbREMRli08hUfXg_uq2PDJqls0XKj6hmtq8pZfGQ8vUJTR8yEkX-1bPel0aT_6qwf2-D1KjLm-JGGCp4XWWLP89C-9jbcFEPbHOMuEbnh5jmXm8tBXp2tcMnFDVWRNvjMU8VtwuRLX4Vr7yqNrIIwSsRlBA9N228fgbTZ81CKg5wxRmC4Emli5YpSDD_TmBj7VVEmpVU9C82Y19FdyeHCBdwoAX6VdK2vrIxZg`

	// K8SAPITimeout 表示超时时间
	K8SAPITimeout = 30
)

var namespace string
var version bool

// GetRootCommand 返回组装好的根命令
func GetRootCommand() *cobra.Command {
	// 定义根命令
	rootCmd := cobra.Command{
		Use: "ks8shell",
		Run: func(cmd *cobra.Command, args []string) {
			if version {
				restclient, err := createK8SClient()
				if err != nil {
					fmt.Println("Err:", err)
					return
				}
				// 通过 ServerVersion 方法来获取版本号
				versionInfo, err := restclient.ServerVersion()
				if err != nil {
					fmt.Println("Err:", err)
					return
				}
				fmt.Println("Kubernetes Version:", versionInfo.String())
			}
		},
	}
	// 添加全局选项参数
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "namespace")

	// 添加显示版本的信息
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "kubernetes version")

	// 添加子命令
	addCommands(&rootCmd)
	return &rootCmd
}

// addCommands 将各个命令拼装在一起
func addCommands(rootCmd *cobra.Command) {
	// Pod
	podRootCmd.AddCommand(&podCreateCmd)
	podRootCmd.AddCommand(&podUpdateCmd)
	podRootCmd.AddCommand(&podGetCmd)
	podRootCmd.AddCommand(&podDeleteCmd)

	// Service
	serviceRootCmd.AddCommand(&serviceCreateCmd)
	serviceRootCmd.AddCommand(&serviceUpdateCmd)
	serviceRootCmd.AddCommand(&serviceGetCmd)
	serviceRootCmd.AddCommand(&serviceDeleteCmd)

	// Ingress
	ingressRootCmd.AddCommand(&ingressCreateCmd)
	ingressRootCmd.AddCommand(&ingressUpdateCmd)
	ingressRootCmd.AddCommand(&ingressGetCmd)
	ingressRootCmd.AddCommand(&ingressDeleteCmd)

	// Secret
	secretRootCmd.AddCommand(&secretCreateCmd)
	secretRootCmd.AddCommand(&secretUpdateCmd)
	secretRootCmd.AddCommand(&secretGetCmd)
	secretRootCmd.AddCommand(&secretDeleteCmd)

	// Deployment
	deploymentRootCmd.AddCommand(&deploymentCreateCmd)
	deploymentRootCmd.AddCommand(&deploymentUpdateCmd)
	deploymentRootCmd.AddCommand(&deploymentGetCmd)
	deploymentRootCmd.AddCommand(&deploymentDeleteCmd)

	// 组装命令
	rootCmd.AddCommand(&podRootCmd)
	rootCmd.AddCommand(&serviceRootCmd)
	rootCmd.AddCommand(&ingressRootCmd)
	rootCmd.AddCommand(&secretRootCmd)
	rootCmd.AddCommand(&deploymentRootCmd)
}

// createK8SClient 根据鉴权信息创建 Kubernetes 的连接客户端
func createK8SClient() (k8sClient *k8s.Clientset, err error) {
	cfg := restclient.Config{}
	cfg.Host = K8SAPIServer
	cfg.CAData = []byte(K8SCertificateData)
	cfg.BearerToken = K8SAPIToken
	cfg.Timeout = time.Second * time.Duration(K8SAPITimeout)
	k8sClient, err = k8s.NewForConfig(&cfg)
	return
}
