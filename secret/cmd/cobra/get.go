package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vitojph/gophercises/secret"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a secret from your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v, err := secret.File(secretsPath(), "my-secret-key")
		if err != nil {
			panic(err)
		}
		key := args[0]
		value, err := v.Get(key)
		if err != nil {
			fmt.Printf("Sorry, no value set for %s.\n", key)
			return
		}
		fmt.Printf("🤫 Secret for %s ==> %s\n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(GetCmd)
}
