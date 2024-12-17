// Instructions

// Write a function that returns true if the string passed as a parameter contains only numerical characters, otherwise, returns false.
// Expected function

// func IsNumeric(s string) bool {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.IsNumeric("010203"))
// 	fmt.Println(piscine.IsNumeric("01,02,03"))
// }

// And its output :

// $ go run .
// true
// false
// $

package main

func IsNumeric(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}
