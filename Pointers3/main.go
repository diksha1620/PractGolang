package main

import "fmt"

func pointer(num *int) {
	*num = 11
	fmt.Println("The adresss:", num)

}

func main() {
	number := 44
	fmt.Println("The value is befor :", number)
	pointer(&number)
	fmt.Println("\nThe value is afer:", number)
}
