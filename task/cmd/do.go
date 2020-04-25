package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as done",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument", arg)
				continue
			}
			ids = append(ids, id)
		}
		fmt.Println("Tasks marked as done:", ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
