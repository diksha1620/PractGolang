package main

import (
	"fmt"
)

var marks = [3]int{11, 22, 11}

func main() {

	fmt.Printf("Array Elements are : %v \n", marks)

	bool()
}

func bool() {
	for i := 0; i < 2; i++ {
		fmt.Println(marks[i])
	}

}
