package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	GetByBrandAndBetweenYears(yearStart int, yearEnd int, brand string) (vehicles []Vehicle, err error)
}
