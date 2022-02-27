package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ~/workspace/go-workspace/gcode/cobra/more-sub-cmd/ [master+*] go run main.go config set-credentials
//this is the set credentials command
// ~/workspace/go-workspace/gcode/cobra/more-sub-cmd/ [master+*]

func main() {
	rootCmd := cobra.Command{
		Use: "kubectl",
	}

	// 一级子命令 config
	configCmd := cobra.Command{
		Use: "config",
	}

	// 二级子命令 set-credentials
	setCredentialsCmd := cobra.Command{
		Use: "set-credentials",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("this is the set credentials command")
		},
	}

	// 组装命令
	configCmd.AddCommand(&setCredentialsCmd)
	rootCmd.AddCommand(&configCmd)

	rootCmd.Execute()
}
