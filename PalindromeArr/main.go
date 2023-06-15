// package main

// import "fmt"

// func main() {
// 	f := func() {

// 		fmt.Println("anonymu/ variable function")

//		}
//		f()
//	}
package main

import (
	"fmt"
)

func main() {
	str := []string{"apple", "bob", "pineapple", "zoz", "tomato", "grapes", "ant", "banana", "boy", "cat", "dog"}
	pStr := invString(str)
	fmt.Println(pStr)
}

// checking the string is palindrome or not
func isPal(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// sending the individual string for checking the palindrome
func invString(str []string) []string {
	result := []string{}
	for _, s := range str {
		if isPal(s) {
			result = append(result, s)
		}
	}
	return result
}
