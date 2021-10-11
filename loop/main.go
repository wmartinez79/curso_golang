package main

import "fmt"

func main() {
	// Conditional For
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
