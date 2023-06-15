package main

import "fmt"

func main() {
	// fmt.Println("Enter the size of array : ")
	// var a int
	// var arr [5]int
	// fmt.Scanln(&a)
	// fmt.Println("Enter the elements of array")
	// for i := 0; i < a; i++ {
	// 	fmt.Scanln(&arr[i])

	// }

	// fmt.Println(arr)

	arr := [...]int{12, 1, 10, 34, 1, 35}
	fmt.Println(arr)
	var large1 = 0
	var large2 = 0
	large1 = arr[0]

	for i := 1; i < len(arr); i++ {
		if large1 < arr[i] {
			large2 = large1
			large1 = arr[i]
		} else if large2 < arr[i] {

			large2 = arr[i]
		}

	}
	fmt.Print("The 2nd largest number is : ", large2)
	fmt.Print("\nThe largest number is : ", large1)

}
