package main

import (
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> b77644cc79295994784c3ce84ab69d5eb1cd3feb
	"log"

	"github.com/vitojph/gophercises/secret"
)

func main() {
<<<<<<< HEAD
	v, err := secret.File(".secrets", "this_is_a_key")
=======
	v, err := secret.File(".secrets", "my-fake-key")
>>>>>>> b77644cc79295994784c3ce84ab69d5eb1cd3feb
	if err != nil {
		panic(err)
	}
	err = v.Set("account1", "my password one")
	log.Println("Set account1")
	if err != nil {
		panic(err)
	}
	err = v.Set("account2", "my password two")
	log.Println("Set account2")
	if err != nil {
		panic(err)
	}
	err = v.Set("account3", "my password three")
	log.Println("Set account3")
	if err != nil {
		panic(err)
	}

	account := "account2"
	plain, err := v.Get(account)
	if err != nil {
		panic(err)
	}
<<<<<<< HEAD
	log.Println("Account:", account, "Plain pass:", plain)

	keys, _ := v.ListSecrets()
	log.Println("Available keys:", keys)

	log.Println("Removing one key")
	v.Remove("account1")

	keys, _ = v.ListSecrets()
	log.Println("Available keys:", keys)
=======
	fmt.Println("Account:", account, "Plain pass:", plain)
>>>>>>> b77644cc79295994784c3ce84ab69d5eb1cd3feb
}
