package main

import "fmt"

type MotorcycleServiceFacade struct {
	engine     *EngineSystem
	brakes     *BrakeSystem
	tires      *TireSystem
	electrical *ElectricalSystem
	report     *ServiceReport
}

func newMotorcycleServiceFacade() *MotorcycleServiceFacade {
	return &MotorcycleServiceFacade{
		engine:     &EngineSystem{},
		brakes:     &BrakeSystem{},
		tires:      &TireSystem{},
		electrical: &ElectricalSystem{},
		report:     &ServiceReport{},
	}
}

func (f *MotorcycleServiceFacade) performService(motorcycleID string) error {
	fmt.Printf("Starting service for motorcycle: %s\n", motorcycleID)

	engineStatus := f.engine.checkEngine()
	brakeStatus := f.brakes.checkBrakes()
	tireStatus := f.tires.checkTires()
	electricalStatus := f.electrical.checkElectrical()

	f.report.generateReport(motorcycleID, engineStatus, brakeStatus, tireStatus, electricalStatus)

	fmt.Println("Service completed successfully")
	return nil
}

