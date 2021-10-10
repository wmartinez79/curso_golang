package main

import "fmt"

func main() {
	// Declaring variables
	helloMessage := "Hello"
	worldMessage := "World"

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s has more than %d cursos\n", nombre, cursos)
	fmt.Printf("%v has more than %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s has more than %d cursos", nombre, cursos)
	fmt.Println(message)

	// data Type
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
