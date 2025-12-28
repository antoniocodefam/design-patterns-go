package main

type Cruiser struct{
	name         string
	engineCC     int
	saddlebags   bool
	forwardControls bool
	chromeLevel  int
}

func (c *Cruiser) inspect(i IInspector){
	i.inspectCruiser(c)
}

func (c *Cruiser) getName() string{
	return c.name
}