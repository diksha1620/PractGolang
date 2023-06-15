package main

import "fmt"

func main() {
	var p = 99

	var r *int = &p
	fmt.Printf("\nthe type p is %T ", p)
	fmt.Printf("\nthe decimal p value is %v: ", p)
	fmt.Printf("\nthe hexadecimal p value is %X: \n", p)

	fmt.Println("\n*****************")

	fmt.Printf("\nthe type r %T ", r)
	fmt.Printf("\nthe decimal value r %v: ", r)
	fmt.Printf("\nthe hexadecimal value  %X: ", r)
}
