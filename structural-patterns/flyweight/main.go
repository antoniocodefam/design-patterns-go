package main

import "fmt"

func main() {
	modelFactory := getMotorcycleModelFactory()
	motorcycleFactory := newMotorcycleFactory()

	hondaForzaModel := modelFactory.getModel("Honda", "Forza")
	hondaForzaModel2 := modelFactory.getModel("Honda", "Forza")
	forza350 := motorcycleFactory.createBike(hondaForzaModel, 350, 2025)
	forza750 := motorcycleFactory.createBike(hondaForzaModel, 750, 2023)
	ymahaTracer := motorcycleFactory.createBike(modelFactory.getModel("Yamaha", "Tracer"), 700, 2024)
	bikes := []IMotorcycle{forza350,forza750,ymahaTracer}

	fmt.Printf("hondaForzaModel and hondaForzaModel2 refer to the same motorcycle instance: %v\n", hondaForzaModel == hondaForzaModel2)
	fmt.Println()

	for i, bike := range bikes {
		model := bike.getModel()
		fmt.Printf("%s %s - %d. %d from %d\n",model.getBrand(),model.getModel(),  i+1, bike.getPower(), bike.getManufactureYear())
	}


}