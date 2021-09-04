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

	// Sum
	x := 10
	y := 50

	result := x + y
	fmt.Println("Sum =", result)

	// Substration
	result = y - x
	fmt.Println("Substration =", result)

	// Multiplication
	result = x * y
	fmt.Println("Multiplication =", result)

	// Division
	result = y / x
	fmt.Println("Division =", result)

	// Mod
	result = y % x
	fmt.Println("Mod =", result)

	// Incremental
	x++
	fmt.Println("Incremental =", x)

	// Decremental
	x--
	fmt.Println("Decremental =", x)

	// Challenge
	// Area Rectangle
	length := 10
	width := 5

	result = length * width
	fmt.Println("Rectangle Area =", result)

	// Area Trapezium
	base1 := 10
	base2 := 5
	height := 3

	result = (base1 * base2 / 2) * height
	fmt.Println("Trapezium Area =", result)

	// Area circle
	radius := 10.00
	var circleArea float64

	circleArea = radius * pi2
	fmt.Println("Cicle Area =", circleArea)

}
