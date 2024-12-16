// Instructions

// Write a function that returns the square root of the int passed as parameter, if that square root is a whole number. Otherwise it returns 0.
// Expected function

// func Sqrt(nb int) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.Sqrt(4))
// 	fmt.Println(piscine.Sqrt(3))
// }

// And its output :

// $ go run .
// 2
// 0
// $

package main

import "fmt"

func Sqrt(nb int) int {
	for i := 1; i < nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}
