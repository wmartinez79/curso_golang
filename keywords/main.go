package main

import "fmt"

func main() {
	//Defer
	defer fmt.Println("Hello")
	fmt.Println("World")

	for i := 0; i < 10; i++ {
		// break
		if i == 8 {
			fmt.Println("Break")
			break
		}

		// continue
		if i == 2 {
			fmt.Println("I am 2")
			continue
		}
		fmt.Println(i)

	}
}
