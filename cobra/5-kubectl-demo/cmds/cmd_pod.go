package cmds

import (
	"github.com/spf13/cobra"
)

// Pod 命令
var podRootCmd = cobra.Command{
	Use:   "pod",
	Short: "pod is used to manage kubernetes Pods",
}

// Pod Create 命令
var podCreateCmd = cobra.Command{
	Use:   "create",
	Short: "create a new pod",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Pod Update 命令
var podUpdateCmd = cobra.Command{
	Use:   "update",
	Short: "update a pod",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Pod Get 命令
var podGetCmd = cobra.Command{
	Use:   "get",
	Short: "get pod or pod list",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Pod Delete 命令
var podDeleteCmd = cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "rm"},
	Short:   "delete a pod",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
