package main

import (
	"fmt"
)

func main() {
	z := []int{1, 2, 3, 4}
	b := z

	copy(b, z)
	fmt.Println(b)
	a := make([]int, 3)
	for i, value := range z {
		a[i] = value
	}

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a[1] = 22
	a[0] = 2
	fmt.Println(a)
	fmt.Printf("arrr len: %d", len(a))

}
