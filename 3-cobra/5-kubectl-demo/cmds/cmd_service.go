package cmds

import (
	"github.com/spf13/cobra"
)

// Service 命令
var serviceRootCmd = cobra.Command{
	Use:     "service",
	Aliases: []string{"svc"},
	Short:   "service is used to manage kubernetes Services",
}

// Service Create 命令
var serviceCreateCmd = cobra.Command{
	Use:   "create",
	Short: "create a new service",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Service Update 命令
var serviceUpdateCmd = cobra.Command{
	Use:   "update",
	Short: "update a service",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Service Get 命令
var serviceGetCmd = cobra.Command{
	Use:   "get",
	Short: "get service or service list",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Service Delete 命令
var serviceDeleteCmd = cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "rm"},
	Short:   "delete a service",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
