package handler

import (
	"app/internal"
	"app/internal/utilities"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
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

		if errors.Is(err, internal.ErrVehicleExistent) {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(ResponseJSON{
				Message: "Identificador del vehículo ya existente.",
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

// FindByColorAndYear returns the list of Vehicles that has that color and year
func (h *VehicleDefault) FindByColorAndYear(w http.ResponseWriter, r *http.Request) {

	// get the path params
	color := chi.URLParam(r, "color")
	year := chi.URLParam(r, "year")

	// parse year to int
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseJSON{
			Message: "Año invalido",
		})
		return
	}

	// call the service
	vehiclesMap, err := h.sv.FindAllEqualTo(internal.EqualFilter{
		Color:           color,
		FabricationYear: yearInt,
	})
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseJSON{
			Message: "Hubo un error interno en el servidor",
		})
		return
	}

	// parse map to slice of VehicleResponseJSON
	vehicles := []VehicleResponseJSON{}
	for _, v := range vehiclesMap {
		// parse model to VehicleResponseJSON
		var vehicleResponse = VehicleResponseJSON{}
		vehicleResponse.parseModelToResponse(v)
		// add to slice
		vehicles = append(vehicles, vehicleResponse)
	}

	if len(vehicles) == 0 {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ResponseJSON{
			Message: "No se encontraron vehiculos con esos criterios",
		})
		return
	}

	// response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseJSON{
		Message: "success",
		Data:    vehicles,
	})

}

// Update updates an existent vehicle
func (h *VehicleDefault) Update(w http.ResponseWriter, r *http.Request) {

	// get id from path param
	id := chi.URLParam(r, "id")

	// parse id to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseJSON{
			Message: "Año invalido",
		})
		return
	}

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
	vehicle.Id = idInt

	// call service
	vehicle, err = h.sv.Update(vehicle)
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

		if errors.Is(err, internal.ErrVehicleNotFound) {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ResponseJSON{
				Message: "No se encontro el vehiculo.",
			})
			return
		}

		if errors.Is(err, internal.ErrVehicleExistent) {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(ResponseJSON{
				Message: "Identificador del vehículo pertenece a otro vehiculo.",
			})
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseJSON{
			Message: "No se pudo actualizar el vehiculo",
		})
		return
	}

	// parse model to response
	var vehicleJSON = VehicleResponseJSON{}
	vehicleJSON.parseModelToResponse(vehicle)

	// response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseJSON{
		Message: "success",
		Data:    vehicleJSON,
	})
}

func (h *VehicleDefault) GetAvgCapacity(w http.ResponseWriter, r *http.Request) {

	// get brand from path param
	brand := chi.URLParam(r, "brand")

	// call the service
	avg, err := h.sv.GetAvgCapacity(brand)
	if err != nil {
		if errors.Is(err, internal.ErrVehiclesNotFound) {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ResponseJSON{
				Message: "No se encontraron vehículos de esa marca.",
			})
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseJSON{
			Message: "Hubo un problema al buscar los vehiculos.",
		})
		return
	}

	// response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseJSON{
		Message: "success",
		Data:    avg,
	})

}
