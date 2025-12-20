package main

type IItem interface{
	print(string)
	clone() IItem
}