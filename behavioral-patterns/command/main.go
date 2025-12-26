package main

func main(){
	ktmDuke := &Motorcycle{}

	powerOnCommand := &PowerOnCommand{
		vehicle: ktmDuke,
	}
	powerOffCommad := &PowerOffCommand{
		vehicle: ktmDuke,
	}

	powerOnBtn := &HandlebarButton{
		command: powerOnCommand,
		id: "On btn",
	}
	powerOffButton := &HandlebarButton{
		command: powerOffCommad,
		id: "Off btn",
	}

	powerOnBtn.press()
    powerOffButton.press()
}