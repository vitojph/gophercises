package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vitojph/gophercises/secret"
)

var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v, err := secret.File(secretsPath(), "my-secret-key")
		if err != nil {
			panic(err)
		}
		key, value := args[0], args[1]
		err = v.Set(key, value)
		if err != nil {
			panic(err)
		}
		fmt.Printf("✨ Secret successfully set for %s! ✨\n", key)

	},
}

func init() {
	RootCmd.AddCommand(SetCmd)
}
