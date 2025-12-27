package main

import "fmt"

type user struct{
	name string
}

func newUser(name string) ISubscriber{
	return &user{
		name: name,
	}
}

func (u *user) notify(brand, model string) {
	fmt.Printf("%s just notified that new %s %s arrived\n", u.name, brand, model)
}

func (u *user) getName() string {
	return u.name
}