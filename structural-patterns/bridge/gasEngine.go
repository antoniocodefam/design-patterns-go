package main

import "fmt"

type GasEngine struct{
}

func (e *GasEngine) startEngine(){
	fmt.Println("Gas engine: vrooom, vrooom!!!")
}

func (e *GasEngine) stopEngine(){
	fmt.Println("Gas engine: engine off")
}