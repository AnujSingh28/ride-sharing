package models

import "github.com/google/uuid"

type Vehicle struct {
	Id            uuid.UUID `json:"id"`
	UserId        uuid.UUID `json:"user_id"`
	VehicleNumber string    `json:"vehicle_number"`
	VehicleModel  string    `json:"vehicle_type"`
}