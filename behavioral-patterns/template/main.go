package main

func main(){
	oilChange := NewOilChangeMaintenance("Yamaha R1", "10W-40 Synthetic")
	oilChange.performMaintenance()

	chainMaint := NewChainMaintenance("Kawasaki Ninja")
	chainMaint.performMaintenance()
}