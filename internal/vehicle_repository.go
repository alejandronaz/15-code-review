package internal

import "errors"

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Add adds a new vehicle to the repo
	Add(newVehicle Vehicle) (v Vehicle, err error)

	// New methods

}

// errors definition
var (
	ErrVehicleExistent = errors.New("vehicle id already exists")
)
