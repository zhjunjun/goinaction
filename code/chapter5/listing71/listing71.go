package main

import (
	"fmt"

	"goinaction/code/chapter5/listing71/entities"
)

func main() {
	u := entities.User{
		Name:  "bill",
		email: "bill@djfek.com",
	}
	//unknown field 'email' in struct literal of type entities.User
	fmt.Printf("User: %v\n", u)
}
