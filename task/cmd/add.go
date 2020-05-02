package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vitojph/gophercises/task/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Ooops, something went wrong:", err.Error())
			return
		}
		fmt.Printf("'%s' added to your TODO list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
