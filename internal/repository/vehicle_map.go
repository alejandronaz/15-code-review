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
	for key := range r.db {
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

	// check if registration already exists
	for _, v := range r.db {
		if v.Registration == newVehicle.Registration {
			return v, internal.ErrVehicleExistent
		}
	}

	// get the last id
	lastId := r.getLastId()
	id := lastId + 1

	// add vehicle
	newVehicle.Id = id
	r.db[id] = newVehicle

	v = r.db[id]
	return
}

// FindAllEqualTo returns a map of vehicles that passed the filters
func (r *VehicleMap) FindAllEqualTo(filter internal.EqualFilter) (v map[int]internal.Vehicle, err error) {

	v = make(map[int]internal.Vehicle)

	/* All of this can be improved using reflect and saving the non-empty fields into a map */
	for key, value := range r.db {
		// if the field is not zero value is because i want to filter using this field

		/* Esto se lee como: si quiero filtrar por este campo, pero el vehiculo no cumple, continuo */
		if filter.Brand != "" && filter.Brand != value.Brand {
			continue
		}

		if filter.Model != "" && filter.Model != value.Model {
			continue
		}

		if filter.Color != "" && filter.Color != value.Color {
			continue
		}

		if filter.FabricationYear != 0 && filter.FabricationYear != value.FabricationYear {
			continue
		}

		if filter.Capacity != 0 && filter.Capacity != value.Capacity {
			continue
		}

		if filter.FuelType != "" && filter.FuelType != value.FuelType {
			continue
		}

		if filter.Transmission != "" && filter.Transmission != value.Transmission {
			continue
		}

		// filters by range

		if filter.FabricationYearRange[0] != 0 && filter.FabricationYearRange[1] != 0 &&
			// is out of range
			(value.FabricationYear < filter.FabricationYearRange[0] || value.FabricationYear > filter.FabricationYearRange[1]) {
			continue
		}

		if filter.LengthRange[0] != 0 && filter.LengthRange[1] != 0 &&
			// is out of range
			(value.Length < filter.LengthRange[0] || value.Length > filter.LengthRange[1]) {
			continue
		}

		if filter.WidthRange[0] != 0 && filter.WidthRange[1] != 0 &&
			// is out of range
			(value.Width < filter.WidthRange[0] || value.Width > filter.WidthRange[1]) {
			continue
		}

		if filter.WeightRange[0] != 0 && filter.WeightRange[1] != 0 &&
			// is out of range
			(value.Weight < filter.WeightRange[0] || value.Weight > filter.WeightRange[1]) {
			continue
		}

		// if all filters passed, add to the map
		v[key] = value
	}

	return v, nil

}

func (r *VehicleMap) Update(vehicle internal.Vehicle) (v internal.Vehicle, err error) {

	// get the vehicle
	v, ok := r.db[vehicle.Id]
	if !ok {
		return v, internal.ErrVehicleNotFound
	}

	// check if registration already exists
	for _, v := range r.db {
		if v.Id != vehicle.Id && v.Registration == vehicle.Registration {
			return v, internal.ErrVehicleExistent
		}
	}

	// update
	v.Brand = vehicle.Brand
	v.Model = vehicle.Model
	v.Registration = vehicle.Registration
	v.Color = vehicle.Color
	v.FabricationYear = vehicle.FabricationYear
	v.Capacity = vehicle.Capacity
	v.MaxSpeed = vehicle.MaxSpeed
	v.FuelType = vehicle.FuelType
	v.Transmission = vehicle.Transmission
	v.Weight = vehicle.Weight
	v.Height = vehicle.Height
	v.Length = vehicle.Length
	v.Width = vehicle.Width

	r.db[vehicle.Id] = v

	return v, nil
}
