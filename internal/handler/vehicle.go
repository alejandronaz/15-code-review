package handler

import (
	"app/internal"
	"app/internal/utilities"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/bootcamp-go/web/response"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		// data := make(map[int]VehicleJSON)
		// for key, value := range v {
		// 	data[key] = VehicleJSON{
		// 		ID:              value.Id,
		// 		Brand:           value.Brand,
		// 		Model:           value.Model,
		// 		Registration:    value.Registration,
		// 		Color:           value.Color,
		// 		FabricationYear: value.FabricationYear,
		// 		Capacity:        value.Capacity,
		// 		MaxSpeed:        value.MaxSpeed,
		// 		FuelType:        value.FuelType,
		// 		Transmission:    value.Transmission,
		// 		Weight:          value.Weight,
		// 		Height:          value.Height,
		// 		Length:          value.Length,
		// 		Width:           value.Width,
		// 	}
		// }

		// return data as array
		data := []VehicleResponseJSON{}
		for _, value := range v {
			newVehicleJSON := VehicleResponseJSON{}
			newVehicleJSON.parseModelToResponse(value)
			data = append(data, newVehicleJSON)
		}

		// response.JSON(w, http.StatusOK, map[string]any{
		// 	"message": "success",
		// 	"data":    data,
		// })

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ResponseJSON{
			Message: "success",
			Data:    data,
		})
	}
}

func (h *VehicleDefault) Add(w http.ResponseWriter, r *http.Request) {

	// get the bytes of body
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseJSON{
			Message: "Datos del vehículo mal formados",
		})
		return
	}

	// deserialize to a map
	bodyMap := make(map[string]any)
	if err := json.Unmarshal(bodyBytes, &bodyMap); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseJSON{
			Message: "Datos del vehículo mal formados",
		})
		return
	}

	// validate if all fields are present
	validFields := utilities.ValidateFields(bodyMap, "brand", "model", "registration", "color", "year", "passengers", "max_speed", "fuel_type", "transmission", "weight", "height", "length", "width")
	if !validFields {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseJSON{
			Message: "Datos del vehículo incompletos",
		})
		return
	}

	// deserialize to a VehicleRequestJSON
	var vehicleReq VehicleRequestJSON
	if err := json.Unmarshal(bodyBytes, &vehicleReq); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseJSON{
			Message: "Datos del vehículo mal formados",
		})
		return
	}

	// parse VehicleRequestJSON to model
	var vehicle internal.Vehicle = vehicleReq.parseRequestToModel()

	// call service
	vehicle, err = h.sv.Add(vehicle)
	if err != nil {
		var target *internal.ErrInvalidAttributes
		if errors.As(err, &target) {
			errInv := err.(*internal.ErrInvalidAttributes)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ResponseJSON{
				Message: fmt.Sprintf("El atributo %s es invalido", errInv.Attr),
			})
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseJSON{
			Message: "Datos del vehículo mal formados",
		})
		return
	}

	// parse Vehicle to VehicleResponse
	data := VehicleResponseJSON{}
	data.parseModelToResponse(vehicle)

	// write response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ResponseJSON{
		Message: "Vehiculo añadido",
		Data:    data,
	})

}
