package handler

import (
	"app/internal"
)

// VehicleRequestJSON is a struct that represents the request body of a vehicle in JSON format
type VehicleRequestJSON struct {
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// parseToModel is a function that parses a vehicle request to a vehicle model
func (req VehicleRequestJSON) parseRequestToModel() internal.Vehicle {
	return internal.Vehicle{
		// Id: 0,
		VehicleAttributes: internal.VehicleAttributes{
			Brand:           req.Brand,
			Model:           req.Model,
			Registration:    req.Registration,
			Color:           req.Color,
			FabricationYear: req.FabricationYear,
			Capacity:        req.Capacity,
			MaxSpeed:        req.MaxSpeed,
			FuelType:        req.FuelType,
			Transmission:    req.Transmission,
			Weight:          req.Weight,
			Dimensions: internal.Dimensions{
				Height: req.Height,
				Length: req.Length,
				Width:  req.Width,
			},
		},
	}
}

// VehicleResponseJSON is a struct that represents the response body of a vehicle in JSON format
type VehicleResponseJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// parseToResponse is a function that parses a vehicle model to a vehicle response
func (res *VehicleResponseJSON) parseModelToResponse(v internal.Vehicle) {
	res.ID = v.Id
	res.Brand = v.Brand
	res.Model = v.Model
	res.Registration = v.Registration
	res.Color = v.Color
	res.FabricationYear = v.FabricationYear
	res.Capacity = v.Capacity
	res.MaxSpeed = v.MaxSpeed
	res.FuelType = v.FuelType
	res.Transmission = v.Transmission
	res.Weight = v.Weight
	res.Height = v.Height
	res.Length = v.Length
	res.Width = v.Width
}

type ResponseJSON struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
