package models

import "github.com/google/uuid"

type User struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age int `json:"age"`
}