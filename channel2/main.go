package main

import (
	"fmt"
)

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Message 1"
	c <- "Message 2"

	fmt.Println(len(c), cap(c))

	// close channel
	close(c)

	// if we try to add another value to a closed channel it will throw an error
	// c <- "Message 3"

	for message := range c {
		fmt.Println(message)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)
	go message("message1", email1)
	go message("message2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Mensaje recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Mensaje recibido de email2", m2)
		}
	}
}
