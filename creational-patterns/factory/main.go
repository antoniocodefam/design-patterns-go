package main

func main() {
	var kia, _ = getCar("Kia")
	var mg, _ = getCar("MG")

	kia.printDetails()
	mg.printDetails()
}