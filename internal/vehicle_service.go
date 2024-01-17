package internal

import (
	"errors"
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
	// Update updates an existent vehicle
	Update(vehicle Vehicle) (v Vehicle, err error)

	// New methods
	// GetAvgCapacity returns the avg of the brands capacity
	GetAvgCapacity(brand string) (avg float64, err error)
}

// errors definition
type ErrInvalidAttributes struct {
	Attr string
}

func (e *ErrInvalidAttributes) Error() string {
	return fmt.Sprintf("atrribute %s is invalid", e.Attr)
}

var (
	ErrVehiclesNotFound = errors.New("vehicles not found")
)
