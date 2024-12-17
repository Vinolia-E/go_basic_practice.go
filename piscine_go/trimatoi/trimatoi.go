// Instructions

//     Write a function that transforms numbers within a string, into an int.

//     If the - sign is encountered before any number it should determine the sign of the returned int.

//     This function should only return an int. In the case of an invalid input, the function should return 0.

//     Note: There will never be more than one sign in a string in the tests.

// Expected function

// func TrimAtoi(s string) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.TrimAtoi("12345"))
// 	fmt.Println(piscine.TrimAtoi("str123ing45"))
// 	fmt.Println(piscine.TrimAtoi("012 345"))
// 	fmt.Println(piscine.TrimAtoi("Hello World!"))
// 	fmt.Println(piscine.TrimAtoi("sd+x1fa2W3s4"))
// 	fmt.Println(piscine.TrimAtoi("sd-x1fa2W3s4"))
// 	fmt.Println(piscine.TrimAtoi("sdx1-fa2W3s4"))
// 	fmt.Println(piscine.TrimAtoi("sdx1+fa2W3s4"))
// }

// And its output :

// $ go run .
// 12345
// 12345
// 12345
// 0
// 1234
// -1234
// 1234
// 1234
// $

package main

import "fmt"

func TrimAtoi(s string) int {
	numstr := ""
	res := 0
	count := 0
	neg := 1

	for _, ch := range s {
		if ch >= '0' && ch <= '9' || ch == '-' {
			if ch == '-' {
				if ch == '-' && numstr != "" {
					continue
				}
				count++
				continue
			}
			numstr += string(ch)
		}
	}
	if count > 1 {
		return 0
	} else if count == 1 {
		neg = -1
	}

	for _, nb := range numstr {
		res = res*10 + (int(nb - '0'))
	}
	return res * neg
}

func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}
