// Instructions

// Write a function that behaves like the Index function.
// Expected function

// func Index(s string, toFind string) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.Index("Hello!", "l"))
// 	fmt.Println(piscine.Index("Salut!", "alu"))
// 	fmt.Println(piscine.Index("Ola!", "hOl"))
// }

// And its output :

// $ go run .
// 2
// 1
// -1
// $

package main

import "fmt"

func Index(s string, toFind string) int {
	for i, ch := range s {
		if ch == rune(toFind[0]) {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}
