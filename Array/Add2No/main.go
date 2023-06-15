package main

import "fmt"

func main() {

	var num1 int
	fmt.Println("Enter 1st Number: ")
	fmt.Scanln(&num1)

	var num2 int
	fmt.Println("Enter 2nd Number: ")
	fmt.Scanln(&num2)

	add := num1 + num2
	fmt.Println("Addition is", add)

	if add/2 == 0 && add > 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("odd Number")
	}

}
