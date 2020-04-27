package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vitojph/gophercises/task/db"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Println("Ooops, something went wrong:", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no pending tasks 🏖")
			return
		}
		fmt.Println("Your TODO has the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
