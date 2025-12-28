package main

type IInspector interface{
	inspectSportBike(s *SportBike)
	inspectCruiser(c *Cruiser)
	inspectAdventureBike(a *AdventureBike)
}