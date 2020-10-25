package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vitojph/gophercises/secret"
)

var RemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes a secret from your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v, err := secret.File(secretsPath(), encodingKey)
		if err != nil {
			panic(err)
		}
		key := args[0]
		res, err := v.Remove(key)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
	},
}

func init() {
	RootCmd.AddCommand(RemoveCmd)
}
