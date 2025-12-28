package main

type IMaintainceProcedure interface {
	performMaintenance()
	prepareBike()
	performSpecificTasks()
	runDiagnostics()
	cleanUp()
}