// Instructions

//     Write a function that counts the runes of a string and that returns that count.

// Expected function

// func StrLen(s string) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	l := piscine.StrLen("Hello World!")
// 	fmt.Println(l)
// }

// And its output :

// $ go run .
// 12
// $

package main

import "fmt"

func StrLen(s string) int {
	count := 0

	for i := 0; i < len(s); i++ {
		count++
	}
	return count
}

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
