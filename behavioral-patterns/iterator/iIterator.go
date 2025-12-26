package main

type IIterator interface{
	getNext() *Motorcycle
	hasMore() bool
}