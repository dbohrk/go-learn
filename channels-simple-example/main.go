package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "ping… \"One ping only\" - Captain Ramius, Red October" }()

	msg := <-messages
	fmt.Println(msg)
}