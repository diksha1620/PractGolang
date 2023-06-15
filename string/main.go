package main

import "fmt"

func main() {

	s := []string{"apple", "bob", "pineapple", "zoz", "tomato", "grapes", "ant", "banana", "boy", "cat", "dog"}
	// string()
	var res []string
	for i:=0;i<len(s);i++{
		palindrome = 
		if s(i)==palindrom
		
	}
	fmt.Println(res)
	fmt.Println(s)

}

func Palindrome(str string) bool {
	lastIdx := len(str) - 1
	// using for loop
	for i := 0; i < lastIdx/2 && i < (lastIdx-i); i++ {
	   if str[i] != str[lastIdx-i] {
		  return false
	   }
	}
}