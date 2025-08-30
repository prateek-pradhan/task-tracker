package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addTaskCli = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task to the task tracker`,
	Run: func(cli *cobra.Command, args []string) {
		fmt.Println(addTask(args[0]))
	},
}

func init() {
	rootCli.AddCommand(addTaskCli)
}