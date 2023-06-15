package main

import "fmt"

func main() {
	fmt.Println("Pattern program Right angle Triangle")
	for i := 0; i <= 4; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
