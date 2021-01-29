package main

import (
	"fmt"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	fmt.Println(quote.Glass())
	// Get a greeting message and print it.
	message := greetings.Hello("Yoshimiki")
	fmt.Println(message)
}
