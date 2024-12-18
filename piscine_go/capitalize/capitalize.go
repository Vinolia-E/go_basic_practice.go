// Instructions

// Write a function that capitalizes the first letter of each word and lowercases the rest.

// A word is a sequence of alphanumeric characters.
// Expected function

// func Capitalize(s string) string {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.Capitalize("Hello! How are you? How+are+things+4you?"))
// }

// And its output :

// $ go run .
// Hello! How Are You? How+Are+Things+4you?
// $

package piscine

func Capitalize(s string) string {
	result := ""
	srune := []rune(s)
	for i := 0; i < len(srune); i++ {
		if i == 0 || srune[i-1] == ' ' || (!num(srune[i-1]) && !Cap(srune[i-1]) && !Low(srune[i-1])) {
			srune[i] = toupper(srune[i])
		} else {
			srune[i] = tolower(srune[i])
		}
		result += string(srune[i])
	}
	return result
}

func toupper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	return r
}

func tolower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

func num(r rune) bool {
	return r >= '0' && r <= '9'
}

func Cap(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func Low(r rune) bool {
	return r >= 'a' && r <= 'z'
}
