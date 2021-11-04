package main

import "fmt"

func main() {
	// you can initialize the value and then pass it to switch or initialize on the switch itself
	// modulo := 5 % 2
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("IS even")
	default:
		fmt.Println("Is odd")
	}

	//switch without condition
	value := 40
	switch {
	case value > 100:
		fmt.Println("Is greater than 100")
	case value < 0:
		fmt.Println("Is lower than 0")
	default:
		fmt.Println("No condition")
	}
}
