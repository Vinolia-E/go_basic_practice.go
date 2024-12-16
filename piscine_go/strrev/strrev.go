// Instructions

//     Write a function that reverses a string.

//     This function will return the reversed string.

// Expected function

// func StrRev(s string) string {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	s := "Hello World!"
// 	s = piscine.StrRev(s)
// 	fmt.Println(s)
// }

// And its output :

// $ go run .
// !dlroW olleH
// $

package main

import "fmt"

func StrRev(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
