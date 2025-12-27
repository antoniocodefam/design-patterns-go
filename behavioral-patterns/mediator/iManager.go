package main

type IManager interface{
	canTruckEnter(truck DeliveryTruck) bool
	truckDeparted(truck DeliveryTruck)
}