package main

import "fmt"

func main() {

	diksha := User{"Diksha", "diksha@gmail.com", true, 21}
	fmt.Printf("Struct is : %+v\n", diksha)
	fmt.Printf("My name is %v and my age is %v", diksha.Name, diksha.Age)
}

type User struct {
	Name   string
	Mail   string
	Status bool
	Age    int
}
