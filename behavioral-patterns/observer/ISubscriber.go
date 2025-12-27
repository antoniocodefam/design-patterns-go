package main

type ISubscriber interface {
	notify(brand, model string)
	getName() string
}