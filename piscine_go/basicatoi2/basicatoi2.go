// Instructions

//     Write a function that simulates the behaviour of the Atoi function in Go. Atoi transforms a number defined as a string in a number defined as an int.

//     Atoi returns 0 if the string is not considered as a valid number. For this exercise non-valid string chains will be tested. Some will contain non-digits characters.

//     For this exercise the handling of the signs + or - does not have to be taken into account.

//     This function will only have to return the int. For this exercise the error return of Atoi is not required.

// Expected function

// func BasicAtoi2(s string) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.BasicAtoi2("12345"))
// 	fmt.Println(piscine.BasicAtoi2("0000000012345"))
// 	fmt.Println(piscine.BasicAtoi2("012 345"))
// 	fmt.Println(piscine.BasicAtoi2("Hello World!"))
// }

// And its output :

// $ go run .
// 12345
// 12345
// 0
// 0
// $

package main

import "fmt"

func BasicAtoi2(s string) int {
	result := 0

	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0
		}
		result = result*10 + (int(ch - '0'))
	}
	return result
}

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}
