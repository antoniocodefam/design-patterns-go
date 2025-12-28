package main

import "fmt"

type BaseMaintenance struct{
	bikeName string
	procedure IMaintainceProcedure
}


func (b *BaseMaintenance) performMaintenance() {
	fmt.Printf("\n=== Maintenance: %s ===\n", b.bikeName)
	b.procedure.prepareBike()
	b.procedure.performSpecificTasks()
	b.procedure.runDiagnostics()
	b.procedure.cleanUp()
	fmt.Println("Maintenance complete")
}

func (b *BaseMaintenance) prepareBike() {
	fmt.Println("Preparing bike: center stand, cool engine, gather tools")
}

func (b *BaseMaintenance) cleanUp() {
	fmt.Println("Cleaning up and logging results")
}
