package main

// NOTE:
// 1. No enums, have to declare a type with constants as values and assignment it to variables
type VehicleType int

const (
	CAR VehicleType = iota
	MOTORCYCLE
)

type Vehicle interface {
	GetLicensePlate() string
	GetType() VehicleType
}

type BaseVehicle struct {
	licensePlate string
	vehicleType  VehicleType
}

func (v *BaseVehicle) GetLicensePlate() string {
	return v.licensePlate
}

func (v *BaseVehicle) GetType() VehicleType {
	return v.vehicleType
}
