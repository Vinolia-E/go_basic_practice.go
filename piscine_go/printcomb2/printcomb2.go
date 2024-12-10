// Write a function that prints in ascending order and on a single line: all possible combinations of two different two-digit numbers.

// These combinations are separated by a comma and a space.
// Expected function

// func PrintComb2() {

// }

package main

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for l := '0'; l <= '9'; l++ {
					if i > k || (i == k && j >= l) {
						continue
					} else if i == '9' && j == '8' && k == '9' && l == '9' {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(l)
						z01.PrintRune('\n')

					} else {

						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(l)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintComb2()
}
