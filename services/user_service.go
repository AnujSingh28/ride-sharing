package services

import (
	"github.com/google/uuid"
	"ride-sharing/models"
)

var (
	AllUsers = map[uuid.UUID]models.User{}
)
func AddUser(name, gender string, age int) (models.User, error) {
	// IsUserAlreadyExists()
	newUser := models.User{
		Id:   uuid.New(),
		Name: name,
		Gender: gender,
		Age: age,
	}
	AllUsers[newUser.Id] = newUser
	return newUser, nil
}


