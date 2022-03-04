package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	msg, err := greetings.Hellos([]string{"Gladys", "Samantha", "Darrin"})
	if err != nil {
		log.Fatal(err)
	}

	for _, msg := range msg {
		fmt.Println(msg)
	}

}
