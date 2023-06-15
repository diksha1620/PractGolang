package main

import "fmt"

func main() {

	diksha := User{"Diksha", "diksha@gmail.com", true, 21}
	fmt.Printf("Struct is : %+v\n", diksha)
	diksha.NewEmail()
	fmt.Printf("My name is %v and my age is %v\n", diksha.Name, diksha.Mail)
	diksha.GetStatus()
}

type User struct {
	Name   string
	Mail   string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("the status is ", u.Status)

}

func (u User) NewEmail() {
	u.Mail = "diksha@go.com"
	fmt.Println("The new email of diksha is ", u.Mail)
}
