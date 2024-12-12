// Instructions

//     Write a function that prints all possible combinations of n different digits in ascending order.

//     n will be defined as : 0 < n < 10

// Below are the references for the printing format expected.

//     (for n = 1) '0, 1, 2, 3, ..., 8, 9'

//     (for n = 3) '012, 013, 014, 015, 016, 017, 018, 019, 023,...689, 789'

// Expected function

// func PrintCombN(n int) {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import "piscine"

// func main() {
// 	piscine.PrintCombN(1)
// 	piscine.PrintCombN(3)
// 	piscine.PrintCombN(9)
// }

// And its output :

// $ go run .
// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
// 012, 013, 014, 015, 016, 017, 018, ... 679, 689, 789
// 012345678, 012345679, ..., 123456789
// $

package main

import (
	"github.com/01-edu/z01"
	"strings"
)

// PrintCombN prints all combinations of n different digits in ascending order
func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	
	finalResult := ""
	result := []string{}
	combination := []byte{}

	backtrack := func(start int) {}

	backtrack = func(start int) {
		if len(combination) == n {
			result = append(result, string(combination))
			return
		}
		for i := start; i <= 9; i++ {
			combination = append(combination, byte(i+'0')) // Convert int to byte (digit character)
			backtrack(i + 1)                               // Move to the next digit
			combination = combination[:len(combination)-1] // Backtrack
		}
	}

	backtrack(0)

	finalResult = strings.Join(result, ", ")

	for _, ch := range finalResult {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func main() {
	PrintCombN(1)
	PrintCombN(2)
	PrintCombN(5)
	PrintCombN(9)
}
