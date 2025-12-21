package main

import "fmt"

type SportBike struct{
	engine IEngine
}

func (b *SportBike) start(){
	fmt.Println("Sport motorcycle starting...")
	b.engine.startEngine()
	fmt.Println("Sport mode activated!")
}

func (b *SportBike) stop(){
	fmt.Println("Sport motorcycle stoping...")
	b.engine.stopEngine()
}

func (b *SportBike) setEngine(engine IEngine){
	b.engine = engine
}