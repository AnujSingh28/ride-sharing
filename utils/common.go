package utils

import "github.com/google/uuid"

func RemoveValueFromSlice(slice []uuid.UUID, value uuid.UUID) []uuid.UUID {
	var result []uuid.UUID
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}