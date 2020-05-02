package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/vitojph/gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "port to start the CYOA web app")
	file := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story from %s.\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JSONStory(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
