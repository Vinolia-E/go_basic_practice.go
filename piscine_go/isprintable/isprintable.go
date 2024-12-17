// Instructions

// Write a function that returns true if the string passed as a parameter contains only printable characters, otherwise, returns false.
// Expected function

// func IsPrintable(s string) bool {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.IsPrintable("Hello"))
// 	fmt.Println(piscine.IsPrintable("Hello\n"))

// }

// And its output :

// $ go run .
// true
// false
// $

package main

import "fmt"

func IsPrintable(s string) bool {
	for _, ch := range s {
		if ch < 32 || ch > 127 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))
}
