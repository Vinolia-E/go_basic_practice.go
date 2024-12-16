// Instructions

//     Write a function that reorders a slice of int in ascending order.

// Expected function

// func SortIntegerTable(table []int) {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	s := []int{5,4,3,2,1,0}
// 	piscine.SortIntegerTable(s)
// 	fmt.Println(s)
// }

// And its output :

// $ go run .
// [0 1 2 3 4 5]
// $

package main

import "fmt"

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := i+1; j < len(table); j++ {
			if table[i] > table[j] {
				table[j],table[i] = table[i],table[j]
			}
		}
	}
}

func main() {
	s := []int{5,4,3,2,1,0}
	SortIntegerTable(s)
	fmt.Println(s)
}
