package main

import (
	"fmt"
	"log"

	"github.com/msanatan/gopractice/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
