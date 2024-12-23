package service

import (
	"app/internal"
	"errors"
)

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

func (s *VehicleDefault) GetByBrandAndBetweenYears(yearStart int, yearEnd int, brand string) (vehicles []internal.Vehicle, err error) {
	dbVehicles, err := s.FindAll()
	if err != nil {
		return nil, errors.New("Internal error to load our databases")
	}

	var filterVehicles []internal.Vehicle
	for _, vehicle := range dbVehicles {
		if vehicle.FabricationYear >= yearStart && vehicle.FabricationYear <= yearEnd && vehicle.Brand == brand {
			filterVehicles = append(filterVehicles, vehicle)
		}
	}

	return filterVehicles, nil
}
