package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCli = &cobra.Command{
	Use: "Task Tracker",
	    Short: "Task Tracker is a cli tool for tracking your tasks",
    Long:  "Task Tracker is a cli tool for tracking your tasks like - Add, Update, Delete and Get.",
    Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error occured while executing Task Tracker '%s'\n", err)
		os.Exit(1)
	}
}