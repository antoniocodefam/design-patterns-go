package main

type IServiceStep interface{
	proceed(vehicle *VehicleService)
	setNext(serviceStep IServiceStep)
}