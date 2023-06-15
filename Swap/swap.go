package main

import "fmt"

func main() {
	a := 10
	b := 12

	a = a + b
	b = a - b
	a = a - b

	fmt.Println(a, b)
}
