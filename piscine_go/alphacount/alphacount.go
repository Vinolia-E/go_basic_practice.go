// Instructions

// Write a function that counts the letters of a string and returns the count.

// The letters are from the Latin alphabet list only, any other characters, symbols or empty spaces shall not be counted.
// Expected function

// func AlphaCount(s string) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	s := "Hello 78 World!    4455 /"
// 	nb := piscine.AlphaCount(s)
// 	fmt.Println(nb)
// }

// And its output :

// $ go run .
// 10
// $

package main

import "fmt"

func AlphaCount(s string) int {
	count := 0

	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			count++
		}
	}
	return count
}

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}
