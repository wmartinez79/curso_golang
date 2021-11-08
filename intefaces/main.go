package main

import "fmt"

type figures2D interface {
	calculateArea() float64
}

type square struct {
	base float64
}

type rectangle struct {
	base   float64
	height float64
}

func calculate(f figures2D) {
	fmt.Println("Area is ", f.calculateArea())
}

func (s square) calculateArea() float64 {
	return s.base * s.base
}

func (r rectangle) calculateArea() float64 {
	return r.base * r.height
}

func main() {
	mySquare := square{base: 10}
	myRectangle := rectangle{base: 5, height: 5}
	calculate(myRectangle)
	calculate(mySquare)

	// interface list
	myInterface := []interface{}{"Hola", 20, 5.12}
	fmt.Println(myInterface...)
}
