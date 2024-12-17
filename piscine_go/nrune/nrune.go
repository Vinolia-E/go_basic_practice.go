// Write a function that returns the nth rune of a string. If not possible, it returns 0.
// Expected function

// func NRune(s string, n int) rune {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"github.com/01-edu/z01"

// 	"piscine"
// )

// func main() {
// 	z01.PrintRune(piscine.NRune("Hello!", 3))
// 	z01.PrintRune(piscine.NRune("Salut!", 2))
// 	z01.PrintRune(piscine.NRune("Bye!", -1))
// 	z01.PrintRune(piscine.NRune("Bye!", 5))
// 	z01.PrintRune(piscine.NRune("Ola!", 4))
// 	z01.PrintRune('\n')
// }

// And its output :

// $ go run .
// la!
// $

package main

import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	if len(s) <= n || n < 0 {
		return 0
	} else if len(s) > n {
		return []rune(s)[n]
	}
	return 0
}

func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}
