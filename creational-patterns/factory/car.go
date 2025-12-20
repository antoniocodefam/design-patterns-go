package main

import "fmt"

type Car struct {
	brand string
	origin string
}

func (c *Car) setBrand(brand string) {
	c.brand = brand
}

func (c *Car) getBrand() string {
	return c.brand
}

func (c *Car) setOrigin(origin string) {
	c.origin = origin
}

func (c *Car) getOrigin() string {
	return c.origin
}

func (c *Car) printDetails() {
	fmt.Printf("Brand: %s - Origin: %s",c.getBrand(),c.getOrigin())
	fmt.Println()
}