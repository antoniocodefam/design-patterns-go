package main

type IComponent interface {
	getPrice() float32
	setPrice(price float32)
	getName() string
	setName(name string)
}

