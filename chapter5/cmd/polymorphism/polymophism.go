package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Println(u.name)
}

type admin struct {
	name  string
	email string
}

func (a *admin) notify() {
	fmt.Println(a.name)
}

func main() {
	bill := user{
		name:  "bill",
		email: "1111",
	}
	sendNotification(&bill)
	admin := admin{
		name:  "admin",
		email: "aaaaa",
	}

	sendNotification(&admin)
}

func sendNotification(n notifier) {
	n.notify()
}
