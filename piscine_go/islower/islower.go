// Instructions

// Write a function that returns true if the string passed as the parameter contains only lowercase characters, otherwise, returns false.
// Expected function

// func IsLower(s string) bool {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.IsLower("hello"))
// 	fmt.Println(piscine.IsLower("hello!"))

// }

// And its output :

// $ go run .
// true
// false
// $

package main

import "fmt"

func IsLower(s string) bool {
	for _, ch := range s {
		if ch < 'a' || ch > 'z' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))
}
