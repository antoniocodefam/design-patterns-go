package main

import "fmt"

type ElectricEngine struct{
}

func (e *ElectricEngine) startEngine(){
	fmt.Println("Electric engine: silent whir!")
}

func (e *ElectricEngine) stopEngine(){
	fmt.Println("Electric engine: power off")
}