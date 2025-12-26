package main

import "fmt"

func main(){
	ktmDuke := &Motorcycle{
		brand: "KTM",
		model: "Duke",
	}
	hondaForza := &Motorcycle{
		brand: "Honda",
		model: "Forza 759",
	}

	motorcycleCollection := &MotorcycleCollection{
		motorcycles: []*Motorcycle{ktmDuke,hondaForza},
	}

	iterator := motorcycleCollection.createIterator()

	for iterator.hasMore() {
		motorcycle := iterator.getNext()
        fmt.Printf("%s - %s\n", motorcycle.brand, motorcycle.model)
	}
}