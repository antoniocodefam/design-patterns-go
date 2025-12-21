package main

import "fmt"

type CruiserBike struct{
	engine IEngine
}

func (b *CruiserBike) start(){
	fmt.Println("Cruiser motorcycle starting...")
	b.engine.startEngine()
	fmt.Println("Comfort mode activated!")
}

func (b *CruiserBike) stop(){
	fmt.Println("Cruiser motorcycle stoping...")
	b.engine.stopEngine()
}

func (b *CruiserBike) setEngine(engine IEngine){
	b.engine = engine
}