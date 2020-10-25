package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vitojph/gophercises/secret"
)

var ListSecretsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists the names of the secrets currently stored",
	Run: func(cmd *cobra.Command, args []string) {
		v, err := secret.File(secretsPath(), encodingKey)
		if err != nil {
			panic(err)
		}
		secrets, _ := v.ListSecrets()
		fmt.Println("Secrets stored 🤫:")
		for _, secret := range secrets {
			fmt.Printf("- %s\n", secret)
		}
	},
}

func init() {
	RootCmd.AddCommand(ListSecretsCmd)
}
