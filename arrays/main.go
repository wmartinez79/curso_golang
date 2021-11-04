package main

import "fmt"

func main() {

	// Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// slice
	slice := []int{0, 5, 3, 2, 4, 6, 7}
	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println("Print first element of the slice ", slice[0])
	fmt.Println("Print first three indexes ", slice[:3])
	fmt.Println("Print range from index 2 to the index lower than 4 ", slice[2:4])
	fmt.Println("Print first three indexes ", slice[4:])

	// Append
	slice = append(slice, 8)
	fmt.Println(slice)

	// Append new list
	newSlice := []int{9, 10, 11}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}
