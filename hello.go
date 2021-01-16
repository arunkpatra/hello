package main

import (
	"fmt"
	"log"

	"github.com/arunkpatra/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message, err := greetings.Hello("Arun")

	// If an error was returned, log it and exit
	if err != nil {
		log.Fatal(err)
	}

	// If no other error was returned, print the message
	fmt.Println(message)
}
