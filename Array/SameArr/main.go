package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 22
	var b string = strconv.Itoa(a)
	fmt.Printf("value : %v , Type: %T\n", b, b)

	arr := [4]int{1, 2, 1, 4}
	fmt.Println(arr)
	var flag int

	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			if arr[i] == arr[j] {
				flag = 1
				break
			}
		}
		if flag == 1 {
			break
		}

	}
	if flag == 1 {
		fmt.Println(flag)
	} else {
		fmt.Println(false)
	}

}
