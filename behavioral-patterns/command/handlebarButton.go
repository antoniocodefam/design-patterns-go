package main

import "fmt"

type HandlebarButton struct{
	id string
	command ICommand
}

func (h *HandlebarButton) press(){
    fmt.Println(h.id + " press")
	h.command.execute()
}