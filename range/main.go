package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) {
	var reversedText string

	text = strings.ToLower(text)
	for i := len(text) - 1; i >= 0; i-- {
		reversedText += string(text[i])
	}

	if reversedText == text {
		fmt.Println("Text ", text, " is Palindrome")
	} else {
		fmt.Println(text, "is not palindrome")
	}
}
func main() {
	slice := []string{"hi", "what", "are", "you", "doing?"}
	for i, valor := range slice {
		fmt.Println(i, valor)
	}
	isPalindrome("Anona")
	isPalindrome("oso")
	isPalindrome("none")

}
