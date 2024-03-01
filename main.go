package main

import (
	"log"
	"ride-sharing/services"
)

func main() {
	rohan,_ := services.AddUser("Rohan", "M", 36)
	services.AddVehicle("KA-01-12345", "Swift", rohan.Id)

	shashank,_ := services.AddUser("Shashank", "M", 29)
	services.AddVehicle("TS-05-62395", "Baleno", shashank.Id)

	services.AddUser("Nandini", "F", 29)

	shipra,_ := services.AddUser("Shipra", "F", 27)
	services.AddVehicle("KA-05-41491", "Polo", shipra.Id)
	services.AddVehicle("KA-12-12332", "Acyiva", shipra.Id)

	services.AddUser("Gaurav", "M", 29)

	rahul,_ := services.AddUser("Rahul", "M", 35)
	services.AddVehicle("KA-05-1234", "XUV", rahul.Id)

	log.Println(services.ShowAllUsers())
	log.Println(services.ShowAllVehicles())

	//services.UserOffersRide()
}
