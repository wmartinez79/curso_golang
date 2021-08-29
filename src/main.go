package main

import "fmt"

func main() {

	// Declaring constants
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("Hola mundo")
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaring int variables

	base := 12
	var altura int = 18
	var area int

	fmt.Println(base, altura, area)

	// Declaring zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Calculating area
	const squareBase = 10
	totalArea := squareBase * squareBase

	fmt.Println("totalArea: ", totalArea)
}
