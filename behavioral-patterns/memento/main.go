package main

func main(){
	motorcycle := Motorcycle{
		model: "Yamaha R1",
		tirePressure: 32.00,
	    absLevel: 3,
		tractionControl: 7,
	}
	motorcycle.displayConfig()

	configManager := newConfigManager()
	configManager.addConfig("street", *motorcycle.createConfig())

	motorcycle.setAbsLevel(1)
	motorcycle.setTirePressure(28)
	motorcycle.setTractionControl(3)
	motorcycle.displayConfig()

	configManager.addConfig("track", *motorcycle.createConfig())
	configManager.listConfigs()

	configManager.loadConfig("tour")
    streetConfig := configManager.loadConfig("street")
   
	motorcycle.setConfig(*streetConfig)
	motorcycle.displayConfig()
}