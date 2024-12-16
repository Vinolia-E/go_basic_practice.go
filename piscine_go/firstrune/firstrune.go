// Instructions

// Write a function that returns the first rune of a string.
// Expected function

// func FirstRune(s string) rune {

// }

package main

import "github.com/01-edu/z01"

func FirstRune(s string) rune {
	for len(s) > 0 {
	  return []rune(s)[0]
	}
	return 0
  }

func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}
