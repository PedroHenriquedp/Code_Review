package handler

import (
	"app/internal"
	"log"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
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
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) GetByBrandAndBetweenYears() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		yearStartParam := chi.URLParam(r, "start_year")
		yearStart, err := strconv.Atoi(yearStartParam)
		if err != nil {
			log.Println(err)
			http.Error(w, "400 Bad Request: É esperado um valor inteiro para o ano inicial", http.StatusBadRequest)
			return
		}

		yearEndParam := chi.URLParam(r, "end_year")
		yearEnd, err := strconv.Atoi(yearEndParam)
		if err != nil {
			log.Println(err)
			http.Error(w, "400 Bad Request: É esperado um valor inteiro para o ano final", http.StatusBadRequest)
			return
		}
		if yearStart <= 0 || yearEnd <= 0 {
			http.Error(w, "400 Bad Request: Os dados dos anos foram mal formatados", http.StatusBadRequest)
			return
		}

		brand := chi.URLParam(r, "brand")
		if brand == "" {
			http.Error(w, "400 Bad Request: Dados mal formatados ou incompletos.", http.StatusBadRequest)
			return
		}

		results, err := h.sv.GetByBrandAndBetweenYears(yearStart, yearEnd, brand)
		if err != nil {
			http.Error(w, "500 - Erro interno", http.StatusBadRequest)
			return
		}

		if results == nil {
			http.Error(w, "404 Not Found: No se encontraron vehículos con esos criterios.", http.StatusNotFound)
			return
		}

		response.JSON(w, http.StatusOK, results)
	}
}
