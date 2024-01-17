package internal

import "errors"

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Add adds a new vehicle to the repo
	Add(newVehicle Vehicle) (v Vehicle, err error)
	// FindAllEqualTo returns a map of vehicles that passed the filters
	FindAllEqualTo(filter EqualFilter) (v map[int]Vehicle, err error)

	// New methods

}

// EqualFilter is a filter for query the repository
type EqualFilter struct {
	// Brand is the brand of the vehicle
	Brand string
	// Model is the model of the vehicle
	Model string
	// Color is the color of the vehicle
	Color string
	// FabricationYear is the fabrication year of the vehicle
	FabricationYear int
	// Capacity is the capacity of people of the vehicle
	Capacity int
	// FuelType is the fuel type of the vehicle
	FuelType string
	// Transmission is the transmission of the vehicle
	Transmission string

	// FabricationYearRange is an array that contains a min and max value for FabricationYear
	FabricationYearRange [2]int
	// LengthRange is an array that contains a min and max value for Length
	LengthRange [2]float64
	// WidthRange is an array that contains a min and max value for Width
	WidthRange [2]float64
	// WeightRange is an array that contains a min and max value for Weight
	WeightRange [2]float64
}

// errors definition
var (
	ErrVehicleExistent = errors.New("vehicle id already exists")
)
