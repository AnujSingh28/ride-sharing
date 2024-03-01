package models

import "github.com/google/uuid"

type Ride struct {
	Id uuid.UUID `json:"id"`
	VehicleId uuid.UUID `json:"vehicle_id"`
	Source string `json:"source"`
	Destination string `json:"destination"`
	SeatCapacity int `json:"seat_capacity"`
	DriverId uuid.UUID `json:"driver_id"`
	PassengerIds []uuid.UUID `json:"passenger_ids"`
	Booked bool `json:"booked"`
}
