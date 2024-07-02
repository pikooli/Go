package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)



func main() {
    // Get a greeting message and print it.
	log.SetPrefix("greetings: ")
    log.SetFlags(0)
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
    // message, err := greetings.Hello("Gladys")
	// message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
    fmt.Println(messages)
}