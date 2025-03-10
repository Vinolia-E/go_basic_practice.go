// Instructions

// Write a recursive function that returns the value of nb to the power of power.

// Negative powers will return 0. Overflows do not have to be dealt with.

// for is forbidden for this exercise.
// Expected function

// func RecursivePower(nb int, power int) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.RecursivePower(4, 3))
// }

// And its output :

// $ go run .
// 64
// $

package main

import "fmt"

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}

func main() {
	fmt.Println(RecursivePower(4, 3))
}
