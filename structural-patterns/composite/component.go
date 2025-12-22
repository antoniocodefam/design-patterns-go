package main

import "fmt"

type Component struct{
	price float32
	name string
}

func (c *Component) setPrice(price float32) {
	c.price = price
}

func (c *Component) getPrice() float32 {
	fmt.Printf("%-40s %10.2f\n", "  " + c.name + "total price", c.price)

	return c.price
}

func (c *Component) getName() string {
	return c.name
}

func (c *Component) setName(name string) {
	c.name = name
}
