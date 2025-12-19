package main

type ICar interface {
	getBrand() string
	getOrigin() string
	setBrand(brand string)
	setOrigin(origin string)
	printDetails()
}