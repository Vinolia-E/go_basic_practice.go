// Instructions

// Write a function that capitalizes each letter in a string.
// Expected function

// func ToUpper(s string) string {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.ToUpper("Hello! How are you?"))
// }

// And its output :

// $ go run .
// HELLO! HOW ARE YOU?
// $

package main

import "fmt"

func ToUpper(s string) string {
	result := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			ch = ch-32
		}
		result += string(ch)
	}
	return result
}

func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}
