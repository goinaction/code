package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type admin struct {
	user person
	id   int
}

func (p *person) Say() {
	fmt.Println(p.age)
	fmt.Println(p.name)
}

func (p person) Change(age int) {
	p.age = age
	fmt.Println(p.age)
	fmt.Println(age)
}
func main() {

	user1 := person{
		age:  22,
		name: "aaa",
	}

	fmt.Println(user1)
	userp := &user1

	fmt.Println()
	user1.Say()
	user1.Change(44)
	userp.Change(55)
	user1.Say()
}
