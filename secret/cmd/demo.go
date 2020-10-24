package main

import (
	"fmt"

	"github.com/vitojph/gophercises/secret"
)

func main() {
	v, err := secret.File(".secrets.txt", "my-fake-key")
	if err != nil {
		panic(err)
	}
	err = v.Set("account1", "my password one")
	if err != nil {
		panic(err)
	}
	err = v.Set("account2", "my password two")
	if err != nil {
		panic(err)
	}
	err = v.Set("account3", "my password three")
	if err != nil {
		panic(err)
	}

	plain, err := v.Get("account2")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)
}
