package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(2)

	message, err := greetings.Hello("Elvis")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
