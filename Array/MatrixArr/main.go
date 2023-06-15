package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 22
	fmt.Println(b)
	var matrix [3][3]int
	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{1, 2, 4}
	matrix[1] = [3]int{1, 2, 4}
	fmt.Println(matrix)

}
