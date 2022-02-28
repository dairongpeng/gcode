package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	podCreateCmd.Flags().StringVar(&Name, "name", "", "pod name")
}

var Name string

// Pod Create 命令
var podCreateCmd = cobra.Command{
	Use:   "create",
	Short: "create a new pod",
	Run: func(cmd *cobra.Command, args []string) {
		if Name == "" {
			cmd.Help()
			return
		}
		fmt.Println("success!")
	},
}
