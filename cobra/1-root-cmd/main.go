package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	// 定义一个命令，直接输出命令行参数
	echoCmd := cobra.Command{
		// 命令名称
		Use: "echo",
		// 命令执行过程
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(strings.Join(args, " "))
		},
	}

	// 执行命令
	echoCmd.Execute()
}
