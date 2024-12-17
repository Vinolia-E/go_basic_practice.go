// Instructions

// Write a function that returns true if the string passed as the parameter only contains alphanumerical characters or is empty, otherwise, and returns false.
// Expected function

// func IsAlpha(s string) bool {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.IsAlpha("Hello! How are you?"))
// 	fmt.Println(piscine.IsAlpha("HelloHowareyou"))
// 	fmt.Println(piscine.IsAlpha("What's this 4?"))
// 	fmt.Println(piscine.IsAlpha("Whatsthis4"))

// }

// And its output :

// $ go run .
// false
// true
// false
// true
// $

package main

import "fmt"

func IsAlpha(s string) bool {
	if s == "" {
		return true
	}
	for _, ch := range s {
		if !islower(ch) && !isupper(ch) && !isnum(ch) {
			return false
		}
	}
	return true
}

func isupper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func islower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isnum(r rune) bool {
	return r >= '0' && r <= '9'
}

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))
}
