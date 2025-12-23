package main

import (
	"fmt"
	"time"
)

type vehicleProduction struct{}

func (p *vehicleProduction) construct(vehicle string){
	fmt.Println("Constructing vehicle")
	time.Sleep(2 * time.Second)
	fmt.Println("Vehicle successfully constructed!")
}