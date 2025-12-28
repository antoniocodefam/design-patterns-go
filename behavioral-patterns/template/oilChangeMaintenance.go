package main

import "fmt"

type oilChangeMaintenance struct {
	BaseMaintenance
	oilType string
}

func NewOilChangeMaintenance(bikeName, oilType string) *oilChangeMaintenance{
	o := &oilChangeMaintenance{
		oilType: oilType,
		BaseMaintenance: BaseMaintenance{
			bikeName: bikeName,
		},
	}

	o.procedure = o

	return o
}


func (o *oilChangeMaintenance) performSpecificTasks() {
	fmt.Println("Oil change: drain old oil, replace filter, add " + o.oilType)
}

func (o *oilChangeMaintenance) runDiagnostics() {
	fmt.Println("Checking for leaks and oil pressure - all good")
}