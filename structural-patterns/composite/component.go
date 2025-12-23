package main

import "fmt"

type component struct{
	price float32
	name string
}

func newComponent(name string, price float32) IComponent {
	return &component{
		name: name,
		price: price,
	}
}

func (c *component) setPrice(price float32) {
	c.price = price
}

func (c *component) getPrice() float32 {
	fmt.Printf("%-40s %10.2f\n", "  " + c.name + "total price", c.price)

	return c.price
}

func (c *component) getName() string {
	return c.name
}

func (c *component) setName(name string) {
	c.name = name
}
