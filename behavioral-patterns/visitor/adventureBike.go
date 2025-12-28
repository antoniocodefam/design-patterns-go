package main

type AdventureBike struct{
	name          string
	engineCC      int
	offRoadTires  bool
	luggage       bool
	groundClearance int
}

func (a *AdventureBike) inspect(i IInspector){
	i.inspectAdventureBike(a)
}

func (a *AdventureBike) getName() string{
	return a.name
}