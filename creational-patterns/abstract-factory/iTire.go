package main

type ITire interface {
	getSize() int
	setSize(size int)
	getType() TireType
	setType(tireType TireType)
}
