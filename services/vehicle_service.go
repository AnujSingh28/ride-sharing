package services

import (
	"github.com/google/uuid"
	"ride-sharing/models"
)

var (
	AllVehicles = map[uuid.UUID]models.Vehicle{}
)

func AddVehicle(vehicleNumber, vehicleModel string, userId uuid.UUID) (models.Vehicle, error) {
	// IsUserAlreadyExists()
	//vehicle type is valid
	newVehicle := models.Vehicle{
		Id:            uuid.New(),
		UserId:        userId,
		VehicleNumber: vehicleNumber,
		VehicleModel:  vehicleModel,
	}
	AllVehicles[newVehicle.Id] = newVehicle
	return newVehicle, nil
}



