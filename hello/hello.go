package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	msg, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}
