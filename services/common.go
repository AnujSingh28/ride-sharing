package services

import (
	"github.com/google/uuid"
	"ride-sharing/models"
)

func ShowAllUsers() map[uuid.UUID]models.User {
	return AllUsers
}

func ShowAllVehicles() map[uuid.UUID]models.Vehicle {
	return AllVehicles
}

func ShowAllAvailableRides() map[uuid.UUID]models.Ride {
	return AllRides
}



