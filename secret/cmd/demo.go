package main

import (
	"log"

	"github.com/vitojph/gophercises/secret"
)

func main() {
	v, err := secret.File(".secrets", "my-fake-key")
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
	log.Println("Account:", account, "Plain pass:", plain)

	keys := v.ListSecrets()
	log.Println("Available keys:", keys)

	log.Println("Removing one key")
	v.Remove("account1")

	keys = v.ListSecrets()
	log.Println("Available keys:", keys)
}
