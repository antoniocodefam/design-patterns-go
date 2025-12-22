package main

import (
	"fmt"
	"slices"
)

type Assembly struct{
	price float32
	name string
	children []IComponent
}

func (a *Assembly) setPrice(price float32) {
	a.price = price
}

func (a *Assembly) getPrice() float32 {
	price := a.price

	for _, child := range(a.children) {
		price += child.getPrice()
	}

   fmt.Printf("%-40s %10.2f\n", a.name + "total price", price)

	return price
}

func (a *Assembly) getName() string {
	return a.name
}

func (a *Assembly) setName(name string) {
	a.name = name
}

func (a *Assembly) add(child IComponent) {
	a.children = append(a.children, child)
}

func (a *Assembly) remove(name string) {
	a.children = slices.DeleteFunc(a.children, func(c IComponent) bool{
		return c.getName() == name
	})
}

func (c *Assembly) clear() {
	c.children = []IComponent{}
}