package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name string
	age  int
}

func (u *user) notify() {
	fmt.Println(u.name)
	fmt.Println(u.age)
}

func main() {
	u := user{
		name: "jack",
		age:  11,
	}
	sendNotification(&u)

}

func sendNotification(n notifier) {
	n.notify()
}
