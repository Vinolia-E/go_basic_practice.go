// Instructions

// Write a function which prints the digits of an int passed in parameter in ascending order. All possible values of type int have to go through, excluding negative numbers. Conversion to int64 is not allowed.
// Expected function

// func PrintNbrInOrder(n int) {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import "piscine"

// func main() {
// 	piscine.PrintNbrInOrder(321)
// 	piscine.PrintNbrInOrder(0)
// 	piscine.PrintNbrInOrder(321)
// }

// And its output :

// $ go run . | cat -e
// 1230123$
// $

package main

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	result := []rune{}
	if n == 0 {
		z01.PrintRune('0')
	}

	for n != 0 {
		mod := n % 10
		result = append(result, rune(mod)+'0')
		n /= 10
	}

	for i := 0; i < len(result);i++ {
		for j := i+1; j < len(result);j++ {
			if result[i] > result [j] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	for _, ch := range result {
		z01.PrintRune(ch)
	}
}

func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}