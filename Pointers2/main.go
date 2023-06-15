package main

import "fmt"

func main() {
	a := 13
	b := &a
	fmt.Println("\nthe value of a : ", a)
	fmt.Println("\nthe adress of a : ", &a)
	fmt.Println("\nthe value of b : ", b)
	fmt.Println("\n the value at b : ", *b)

}
