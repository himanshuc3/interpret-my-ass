package main

// NOTE:
// 1. Run this package using go run folder/name/*.go
func main() {
	parkingLot := GetParkingLotInstance()
	parkingLot.AddLevel(NewLevel(1, 33))
	parkingLot.AddLevel(NewLevel(2, 99))

	car := NewCar("CCC333")
	bike := NewBike("DYD3434")

	parkingLot.ParkVehicle(car)
	parkingLot.ParkVehicle(bike)

	// parkingLot.DisplayAvailability()
	parkingLot.UnparkVehicle(bike)
	parkingLot.DisplayAvailability()
}
