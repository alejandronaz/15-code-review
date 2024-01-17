package repository

import "app/internal"

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// getLastId is a method that returns the last id of the db
func (r *VehicleMap) getLastId() (id int) {
	for key, _ := range r.db {
		if key > id {
			id = key
		}
	}
	return
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// Add is a method that adds a new vehicle to the db
func (r *VehicleMap) Add(newVehicle internal.Vehicle) (v internal.Vehicle, err error) {
	// get the last id
	lastId := r.getLastId()
	id := lastId + 1

	// add vehicle
	newVehicle.Id = id
	r.db[id] = newVehicle

	v = r.db[id]
	return
}
