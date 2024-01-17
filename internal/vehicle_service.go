package internal

import (
	"fmt"
)

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Add adds a new vehicle to the repo
	Add(newVehicle Vehicle) (v Vehicle, err error)
	// FindAllEqualTo returns a map of vehicles that passed the filters
	FindAllEqualTo(filter EqualFilter) (v map[int]Vehicle, err error)

	// New methods

}

// errors definition
type ErrInvalidAttributes struct {
	Attr string
}

func (e *ErrInvalidAttributes) Error() string {
	return fmt.Sprintf("atrribute %s is invalid", e.Attr)
}
