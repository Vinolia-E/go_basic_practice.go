// Write a function that prints an int passed in parameter. All possible values of type int have to go through. You cannot convert to int64.
// Expected function

// func PrintNbr(n int) {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"piscine"
// 	"github.com/01-edu/z01"
// )
// func main() {
// 	piscine.PrintNbr(-123)
// 	piscine.PrintNbr(0)
// 	piscine.PrintNbr(123)
// 	z01.PrintRune('\n')
// }

// And its output :

// $ go run .
// -1230123
// $

package main

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	result := ""
	neg := ""

	if n == 0 {
		result += "0"
	}

	if n < 0 {
		neg = "-"
		// n *= - 1
		for n != 0 {
			mod := (n % 10) * -1
			result = string(mod+'0') + result
			n /= 10
		}
	}
	for n > 0 {
		mod := n % 10
		result = string(mod+'0') + result
		n /= 10
	}
	result = neg + result
	for _, ch := range result {
		z01.PrintRune(ch)
	}
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	PrintNbr(-9223372036854775808)
	z01.PrintRune('\n')
}
