package services

import (
	"errors"
	"github.com/google/uuid"
	"ride-sharing/models"
	"ride-sharing/utils"
	"sort"
)

var (
	AllRides = map[uuid.UUID]models.Ride{}
	BookedRides = map[uuid.UUID]models.Ride{}
)

func UserOffersRide(vehicleId, userId uuid.UUID, source, dest string, seats int) (models.Ride, error) {
	newRide := models.Ride{
		Id:           uuid.New(),
		VehicleId:    vehicleId,
		Source:       source,
		Destination:  dest,
		SeatCapacity: seats,
		DriverId:     userId,
		PassengerIds: make([]uuid.UUID,0),
		Booked:       false,
	}
	AllRides[newRide.Id] = newRide
	return newRide, nil
}

func PassengerBooksRide( userId uuid.UUID, source, dest, vehicleType string, seatsRequired int) (models.Ride, error){
	var potentialRides []models.Ride
	type rideSeat struct {
		RideId uuid.UUID
		Seats  int
	}
	var vacantSeats []rideSeat

	for rideId, ride := range AllRides {
		if ride.Source == source && ride.Destination == dest && ride.SeatCapacity >= seatsRequired + len(ride.PassengerIds) {
			potentialRides = append(potentialRides, ride)
			vacantSeats = append(vacantSeats, rideSeat{RideId: rideId, Seats: ride.SeatCapacity - len(ride.PassengerIds)})
		}
	}

	if len(potentialRides) == 0 {
		return models.Ride{}, errors.New("No ride exist for given source and destination ")
	}

	if vehicleType == "" {
		for _, ride := range potentialRides {
			if AllVehicles[ride.VehicleId].VehicleModel == vehicleType {
				ride.PassengerIds = append(ride.PassengerIds, userId)
				ride.Booked = true
				return ride, nil
			}
		}
		return models.Ride{}, errors.New("No ride exist for: " + vehicleType)
	} else {
		sort.SliceStable(vacantSeats, func(i, j int) bool {
			return vacantSeats[i].Seats > vacantSeats[j].Seats
		})
		bookedRide := AllRides[vacantSeats[0].RideId]
		bookedRide.PassengerIds = append(bookedRide.PassengerIds, userId)
		bookedRide.Booked = true
		return bookedRide, nil
	}
}

func PassengerCancelsRide(userId, rideId uuid.UUID) models.Ride {
	ride := AllRides[rideId]
	passengers := utils.RemoveValueFromSlice(ride.PassengerIds, userId)
	ride.PassengerIds = passengers
	if len(passengers) == 0 {
		ride.Booked = false
	}
	AllRides[rideId] = ride
	return ride
}
