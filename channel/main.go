package main

import "fmt"

// channel c will be use just as an input channel <-
func say(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("bye", c)

	fmt.Println(<-c)

}
