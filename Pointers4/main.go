package main

import "fmt"

func main() {
	display()
	fmt.Println(display())

}

func display() *string {
	msg := "Hello"
	return &msg
}
