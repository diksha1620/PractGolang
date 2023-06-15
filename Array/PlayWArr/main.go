package main

import (
	"fmt"
)

func main() {
	var arr = [...]int{1, 2, 3, 4}
	fmt.Println("Initial Araay is :", arr)
	arr[0] = 11
	arr[2] = 12
	fmt.Println("chaning element Araay is :", arr)
	b := &arr
	b[1] = 10
	fmt.Println("Pointer Array is: ", b)
	var matArr [3][3]int
	matArr[0] = [3]int{1, 2, 3}
	matArr[1] = [3]int{1, 3}
	// var matArr [4][4]int
	// matArr[0] = [4]int{0, 1, 2}
	// matArr[3] = [4]int{1, 2, 2}
	// matArr[1] = [4]int{1, 1, 2, 99}
	fmt.Println("Matrix Array is: ", matArr)

	name := [...]string{"Happy", "Happy3"}
	fmt.Println("Length of Array is: ", len(name[1]))
	fmt.Println("capacity of Array is: ", cap(name))
}
