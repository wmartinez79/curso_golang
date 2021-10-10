package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (b, c int) {
	return a, a * 2
}

func main() {
	normalFunction("hello world!")
	tripleArgument(1, 3, "bye")

	newVal := returnValue(2)
	fmt.Println("Value:", newVal)

	value1, value2 := doubleReturn(2)
	fmt.Println("value 1 =", value1, "value 2 = ", value2)

	value3, _ := doubleReturn(4)
	fmt.Println("value 3 =", value3)
}
