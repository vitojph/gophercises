package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vitojph/gophercises/task/db"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as done",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Println("Ooops, something went wrong when fetching tasks", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task ID:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark '%d' as completed: %s\n", id, err)
			} else {
				fmt.Printf("Marked '%d' as completed.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
