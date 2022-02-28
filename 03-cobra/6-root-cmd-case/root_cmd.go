package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	podCreateCmd.Flags().StringVar(&Name, "name", "", "pod name")
}

var Name string

var podCreateCmd = cobra.Command{
	Use:   "show-name",
	Short: "show name",
	Run: func(cmd *cobra.Command, args []string) {
		if Name == "" {
			cmd.Help()
			return
		}
		fmt.Println(Name)
	},
}
