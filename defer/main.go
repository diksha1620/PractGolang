package main

import "fmt"

func main() {
	var i int
	for i = 0; i < 12; i++ {
		if i == 6 {
			panic("Not working")

		}
		fmt.Println(i)
	}
	defer fmt.Println("Not true")

	fmt.Println(i)

}
