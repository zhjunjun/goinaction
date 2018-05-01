// Sample program to show how to declare methods and how the Go
// compiler supports them.
package main

import (
	"fmt"
)

// type Duration int64

// user defines a user in the program.
type user struct {
	name  string
	email string
}

type admin struct {
	person user
	level  string
}

func (a admin) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		a.person.name,
		a.person.email)
	a.person.name = "zhangjunjun"
}

func (a *admin) notifyPtr() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		a.person.name,
		a.person.email)
	a.person.name = "zhangjunjun"
}

// notify implements a method with a value receiver.
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	u.email = email
}

func init() {
	fmt.Println("hello ......")
}

// main is the entry point for the application.
func main() {
	// Values of type user can be used to call methods
	// declared with a value receiver.
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	// Pointers of type user can also be used to call methods
	// declared with a value receiver.
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	// Values of type user can be used to call methods
	// declared with a pointer receiver.
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// Pointers of type user can be used to call methods
	// declared with a pointer receiver.
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()

	master := admin{user{"zjj", "ycuzjj@163.com"}, "master"}
	master.notify()
	master.notify()

	master.notifyPtr()
	master.notifyPtr()
	// var dur Duration
	// dur = int64(10000)
}
