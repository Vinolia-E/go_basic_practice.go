// Instructions

// Write an iterative function that returns the factorial of the int passed as parameter.

// Errors (non possible values or overflows) will return 0.
// Expected function

// func IterativeFactorial(nb int) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	arg := 4
// 	fmt.Println(piscine.IterativeFactorial(arg))
// }

// And its output :

// $ go run .
// 24
// $

package main

import "fmt"

func IterativeFactorial(nb int) int {
	result := 0
	if nb == 0 {
		return 1
	}

	for i := 1; i < nb; i++ {
		result += nb * i
		// fmt.Println(result)
	}
	return result
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
