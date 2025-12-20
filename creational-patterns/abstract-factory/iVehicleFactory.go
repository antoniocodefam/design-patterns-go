package main

import "fmt"

type IVehicleFacotry interface {
	createMotorcycle() IMotorcycle
	createTire() ITire
}

func GetVehicleFactory(motorcycleType string) (IVehicleFacotry,error) {
	if motorcycleType == "dirt" {
		return &DirtMotorcycleFactory{}, nil
	}

	if motorcycleType == "city" {
		return  &CityMotorcycleFactory{}, nil
	}

    return nil, fmt.Errorf("Wrong motorcycle type")
}