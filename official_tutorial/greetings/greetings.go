package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return anerror with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// Create a message using a rondom format.
	message := fmt.Sprintf(RandomFormat(), name)
	return message, nil
}

/*
	Init sets initial values for variables used in the function.
	Go executes init functions automatically at program startup,
	after global variables have been initialized.
*/
func init() {
	rand.Seed(time.Now().UnixNano())
}

/*
	RandomFormat returns one of a set of greeting messages. The returned
	message is selected at random.
*/
func RandomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	/*
		Return a randomely selected message format by specifying
		a random index for the slice of formats.
	*/
	return formats[rand.Intn(len(formats))]
}
