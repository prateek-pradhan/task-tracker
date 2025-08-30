package cli

import (
	"github.com/spf13/cobra"
)

var listTasksCli = &cobra.Command{
	Use:   "list",
	Short: "list tasks",
	Long:  `list tasks from task tracker tracker`,
	Run: func(cli *cobra.Command, args []string) {
		var status string
		if len(args) > 0 { 
			status = args[0]
		}
		list(status)
	},
}

func init() {
	rootCli.AddCommand(listTasksCli)
}