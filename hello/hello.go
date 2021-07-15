package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	msg, err := greetings.Hello("flat35hd99")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}
