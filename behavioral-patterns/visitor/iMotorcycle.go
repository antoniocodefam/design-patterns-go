package main

type IMotorcycle interface{
	inspect(i IInspector)
	getName() string
}