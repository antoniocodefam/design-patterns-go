package main

func main(){
	manager := &DeliveryManager{}

	truck1 := &DeliveryTruck{
		id: 1,
		manager: manager,
	}
	truck2 := &DeliveryTruck{
		id: 2,
		manager: manager,
	}
	truck3 := &DeliveryTruck{
		id: 3,
		manager: manager,
	}

	truck1.arrive()
	truck2.arrive()
	truck3.arrive()

	truck1.depart()
	truck3.depart()
}