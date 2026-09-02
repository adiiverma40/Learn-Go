package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structus")

	u := User{
		Name:   "John",
		Email:  "john@example.com",
		Status: false,
		Age:    0,
	}
	fmt.Println(u)
	fmt.Printf("%+v \n", u)
	fmt.Printf("%v \n", u.Name)
	u.GetStatus()
	u.NewMail()
	fmt.Println(u.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int // not exportable
}

// the copy of the user struct is passing not the actual struct 
func (u User) GetStatus(){
	fmt.Println("Is user active : ", u.Status)
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("New Email is: ", u.Email)
}
// to change the actual email. we have to pass the pointer 