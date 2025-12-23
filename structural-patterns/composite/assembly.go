package main

import (
	"fmt"
	"slices"
)

type assembly struct{
	price float32
	name string
	children []IComponent
}

func newAssembly(name string) IComposite {
	return &assembly{
		name: name,
		children: []IComponent{},
	}
}

func (a *assembly) setPrice(price float32) {
	a.price = price
}

func (a *assembly) getPrice() float32 {
	price := a.price

	for _, child := range(a.children) {
		price += child.getPrice()
	}

   fmt.Printf("%-40s %10.2f\n", a.name + "total price", price)

	return price
}

func (a *assembly) getName() string {
	return a.name
}

func (a *assembly) setName(name string) {
	a.name = name
}

func (a *assembly) add(child IComponent) {
	a.children = append(a.children, child)
}

func (a *assembly) remove(name string) {
	a.children = slices.DeleteFunc(a.children, func(c IComponent) bool{
		return c.getName() == name
	})
}

func (c *assembly) clear() {
	c.children = []IComponent{}
}