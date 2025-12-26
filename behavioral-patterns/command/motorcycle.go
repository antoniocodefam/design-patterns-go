package main

import "fmt"

type Motorcycle struct{
	isRunning bool
}

func (m *Motorcycle) on(){
	fmt.Println("Motorcycle engine on!")
	m.isRunning = true
}

func (m *Motorcycle) off(){
	fmt.Println("Motorcycle engine off!")
	m.isRunning = false
}