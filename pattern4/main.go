package main

import "fmt"

func main() {
	fmt.Println("Pattern problem for Numbers")
	n := 5
	for i := 1; i < n; i++ {
		for j := i; j < n-1; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= (2*i - 1); j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
