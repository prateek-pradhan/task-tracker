package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var markInProgressTaskCli = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "Update status to In Progress",
	Long:  `Update status to In Progress of task task tracker`,
	Run: func(cli *cobra.Command, args []string) {
		fmt.Println(updateTaskStatus(args[0], "In Progress"))
	},
}

var markDoneTaskCli = &cobra.Command{
	Use:   "mark-done",
	Short: "Update status to Done",
	Long:  `Update status to Done of task task tracker`,
	Run: func(cli *cobra.Command, args []string) {
		fmt.Println(updateTaskStatus(args[0], "Done"))
	},
}

func init(){
	rootCli.AddCommand(markInProgressTaskCli, markDoneTaskCli)
}