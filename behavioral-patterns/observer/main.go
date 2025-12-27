package main

func main(){
	john, pete := newUser("John"), newUser("Pete")
	motorcycleFactory := newMotorcycleFactory()

	motorcycleFactory.subscribe(john, "Yamaha")
	motorcycleFactory.subscribe(pete, "Honda")
	motorcycleFactory.subscribe(pete, "Yamaha")

	motorcycleFactory.createMotorcycle("Yamaha", "R1")
	motorcycleFactory.createMotorcycle("Honda", "CBR 500")

	motorcycleFactory.unsubscribe(pete, "Yamaha")

	motorcycleFactory.createMotorcycle("Yamaha", "MT-07")
}