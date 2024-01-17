package service

import "app/internal"

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

// Add is a method that adds a new vehicle
func (s *VehicleDefault) Add(newVehicle internal.Vehicle) (v internal.Vehicle, err error) {

	// check if the fabrication year is valid
	if newVehicle.FabricationYear < 1886 {
		return v, &internal.ErrInvalidAttributes{Attr: "FabricationYear"}
	}

	// some others validations
	// ...

	// add the vehicle
	v, err = s.rp.Add(newVehicle)
	if err != nil {
		return internal.Vehicle{}, err
	}

	return
}

// FindAllEqualTo returns a map of vehicles that passed the filters
func (s *VehicleDefault) FindAllEqualTo(filter internal.EqualFilter) (v map[int]internal.Vehicle, err error) {
	// call the repo
	v, err = s.rp.FindAllEqualTo(filter)
	return
}

// Update updates an existent vehicle
func (s *VehicleDefault) Update(vehicle internal.Vehicle) (v internal.Vehicle, err error) {
	// check if the fabrication year is valid
	if vehicle.FabricationYear < 1886 {
		return v, &internal.ErrInvalidAttributes{Attr: "FabricationYear"}
	}

	// some others validations
	// ...

	// call the repo
	v, err = s.rp.Update(vehicle)
	if err != nil {
		return v, err
	}

	return
}
