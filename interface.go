package main

import (
	"fmt"
)

type User interface {
	PrintName()
	PrintDetails()
}

type Person struct {
	name, email string
}

func (p Person) PrintName() {
	fmt.Println(p.name)
}

func (p Person) PrintDetails() {
	fmt.Println(p.name, p.email)
}

func (p *Person) changeName(newName string) {
	p.name = newName
}

type Admin struct {
	name, email string
}

func (p Admin) PrintName() {
	fmt.Println(p.name)
}

func (p Admin) PrintDetails() {
	fmt.Println(p.name, p.email)
}

type Guest struct {
	name  string
	Users []User
}

func (g Guest) PrintDetails() {
	fmt.Println(g.name)
	for _, u := range g.Users {
		u.PrintName()
		u.PrintDetails()
	}
}

func main() {
	person := Person{"John", "john@mail.com"}
	person.PrintName()
	ch := &person
	ch.changeName("Just")
	admin := Admin{"Jane", "jane@jane"}
	guest := Guest{"jame", []User{person, admin}}
	guest.PrintDetails()

}
