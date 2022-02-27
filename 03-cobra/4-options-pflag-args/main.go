package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ~/workspace/go-workspace/gcode/cobra/4-options-pflag-args/ [master+*] go run main.go -n devops get -o json -l app=echo-go pods echo-go-67986c999b-vbrtl
//Print flags...
//Flags: namespace=[devops], selector=[app=echo-go], output=[json]
//Print args...
//Arg: pods
//Arg: echo-go-67986c999b-vbrtl
// ~/workspace/go-workspace/gcode/cobra/4-options-pflag-args/ [master+*]

var namespace string

// 在命令行应用中除了选项参数之外，剩下的参数都可以通过 Args 来获取
func main() {
	rootCmd := cobra.Command{
		Use: "kubectl",
	}
	// 添加全局的选项，所有的子命令都可以继承
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "If present, the namespace scope for this CLI request")

	// 一级子命令 get
	var outputFormat string
	var labelSelector string
	getCmd := cobra.Command{
		Use: "get",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print flags...")
			fmt.Printf("Flags: namespace=[%s], selector=[%s], output=[%s]\n", namespace, labelSelector, outputFormat)
			fmt.Println("Print args...")
			for _, arg := range args {
				fmt.Println("Arg:", arg)
			}
		},
	}
	// 添加命令选项，这些选项仅 get 命令可用
	getCmd.Flags().StringVarP(&outputFormat, "output", "o", "", "Output format")
	getCmd.Flags().StringVarP(&labelSelector, "selector", "l", "", "Selector (label query) to filter on")

	// 组装命令
	rootCmd.AddCommand(&getCmd)

	// 执行命令
	rootCmd.Execute()
}
