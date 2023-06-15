package main

import "fmt"

func main() {
	var a int

	fmt.Printf("Enter the number: ")
	fmt.Scanln(&a)

	rev := 0
	for a != 0 {
		rem := a % 10
		rev = rev*10 + rem
		a = a / 10
	}
	fmt.Println(rev)

	// rev string
	str := "Madam"

	var result string
	for _, v := range str {
		result = string(v) + result
	}
	fmt.Println(result)

}
