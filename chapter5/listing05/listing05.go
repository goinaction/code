// Sample program to show how polymorphic behavior with interfaces.
package main

import (
	"fmt"
)

// notifier is an interface that defines notification
// type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// admin defines a admin in the program.
type admin struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending user Email To %s<%s>\n",
		u.name,
		u.email)
}

// notify implements the notifier interface with a pointer receiver.
func (a *admin) notify() {
	fmt.Printf("Sending admin Email To %s<%s>\n",
		a.name,
		a.email)
}

// main is the entry point for the application.
func main() {
	// Create a user value and pass it to sendNotification.
	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)

	// Create an admin value and pass it to sendNotification.
	jill := admin{"Jill", "jill@email.com"}
	sendNotification(&jill)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(notify notifier) {
	notify.notify()
}
