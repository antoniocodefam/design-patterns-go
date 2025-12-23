package main

import (
	"fmt"
	"time"
)

type vehicleProductionManager struct {
	vehicleProduction IVehicleProduction
	constructedVehicles int
}

const SERVICE_REQUIRED_LIMIT = 3

var manager = &vehicleProductionManager{
	constructedVehicles: 0,
	vehicleProduction: &vehicleProduction{},
}
func (m *vehicleProductionManager) construct(vehicle string) {
	if m.constructedVehicles >= SERVICE_REQUIRED_LIMIT {
		fmt.Println("Service required! Performing service...")
		time.Sleep(5 * time.Second)
		fmt.Println("Service finished, ready to go!")
		m.constructedVehicles = 0
	}

	m.vehicleProduction.construct(vehicle)
	m.constructedVehicles++
}

func getProductionManager() IVehicleProduction {
	return manager
}