package validation

import "github.com/google/uuid"

func ValidateID(id string) error {
	return uuid.Validate(id)
}
