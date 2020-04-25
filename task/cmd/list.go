package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Your TODO list.")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
