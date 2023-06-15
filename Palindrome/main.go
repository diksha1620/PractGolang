package main

import "fmt"

func main() {
	a := isPalindrome(121121)

	fmt.Println(a)
}
func isPalindrome(x int) bool {
	reversed := 0
	original := x
	if x < 0 {
		return false
	}
	for x > 0 {
		remainder := x % 10
		reversed = (reversed * 10) + remainder
		x = x / 10
	}
	return original == reversed

}
