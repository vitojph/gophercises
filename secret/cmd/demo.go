package main

import (
	"fmt"

	"github.com/vitojph/gophercises/secret"
)

func main() {
	v := secret.Memory("my-fake-key")
	err := v.Set("twitter", "my twitter password")
	if err != nil {
		panic(err)
	}

	plain, err := v.Get("twitter")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)
}
