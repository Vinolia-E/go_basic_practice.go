// Instructions

// Write a function that lower cases for each letter in a string.
// Expected function

// func ToLower(s string) string {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.ToLower("Hello! How are you?"))
// }

// And its output :

// $ go run .
// hello! how are you?
// $

package main

import "fmt"

func ToLower(s string) string {
	result := ""
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			ch += 32
		}
		result += string(ch)
	}
	return result
}

func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}
