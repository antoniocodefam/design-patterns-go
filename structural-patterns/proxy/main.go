package main

import "fmt"

func main(){
	productionManager := getProductionManager()
	vehicles := []string{"Honda", "Yamaha", "CF Moto", "Aprillia"}

	for i, vehicle := range vehicles {
		fmt.Printf("%d.Requesting new %s\n", i+1, vehicle)
		productionManager.construct(vehicle)
	}
}