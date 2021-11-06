package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Joe"] = 21
	m["Will"] = 42

	fmt.Println(m)

	// Iterate map
	for index, value := range m {
		fmt.Println(index, value)
	}

	// Find if the key exists in the map
	// no exists
	value, ok := m["None"]
	fmt.Println(value, ok)
	// bingo
	value2, ok2 := m["Will"]
	fmt.Println(value2, ok2)
}
