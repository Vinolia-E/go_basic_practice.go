// Instructions

// Write a function that returns the last rune of a string.
// Expected function

// func LastRune(s string) rune {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"github.com/01-edu/z01"

// 	"piscine"
// )

// func main() {
// 	z01.PrintRune(piscine.LastRune("Hello!"))
// 	z01.PrintRune(piscine.LastRune("Salut!"))
// 	z01.PrintRune(piscine.LastRune("Ola!"))
// 	z01.PrintRune('\n')
// }

// And its output :

// $ go run .
// !!!
// $

package main

func LastRune(s string) rune {
	for len(s) > 0 {
		return []rune(s)[len(s)-1]
	}
	return 0
}