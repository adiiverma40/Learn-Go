// IMPORTANT
// NO INHEARITNCE,
// NO SUPER OR PARENT OR CHILD
//

package main

import "fmt"

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

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
