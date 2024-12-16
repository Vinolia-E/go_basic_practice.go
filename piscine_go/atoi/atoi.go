// Write a function that simulates the behaviour of the Atoi function in Go. Atoi transforms a number represented as a string in a number represented as an int.

// Atoi returns 0 if the string is not considered as a valid number. For this exercise non-valid string chains will be tested. Some will contain non-digits characters.

// For this exercise the handling of the signs + or - does have to be taken into account.

// This function will only have to return the int. For this exercise the error result of Atoi is not required.

// Expected function

// func Atoi(s string) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// "fmt"
// "piscine"
// )

// func main() {
// fmt.Println(piscine.Atoi("12345"))
// fmt.Println(piscine.Atoi("0000000012345"))
// fmt.Println(piscine.Atoi("012 345"))
// fmt.Println(piscine.Atoi("Hello World!"))
// fmt.Println(piscine.Atoi("+1234"))
// fmt.Println(piscine.Atoi("-1234"))
// fmt.Println(piscine.Atoi("++1234"))
// fmt.Println(piscine.Atoi("--1234"))
// }

// And its output :

// $ go run .
// 12345
// 12345
// 0
// 0
// 1234
// -1234
// 0
// 0
// $

package main

func Atoi(s string) int {
  result := 0
  neg := 1

  for i, ch := range s {
    if ch < '0' || ch > '9' {
      if ch == '+' && i == 0 {
        s = s[1:]
         continue
      } else if ch == '-' && i == 0 {
        neg = -1
        s = s[1:]
         continue
      } else {
        return 0
      }
    }
    result = result*10 + int(ch-'0')
  }
  return result*neg
}
