package kernel

import (
	"autopark-service/internal/autopark/car"
	"autopark-service/internal/db"
	"autopark-service/internal/std"
)

// Create car store car in storage
//
// Pre-cond: given car entity and entity implementing CarStorer interafce
//
// Post-cond: executes command to store given car in given CarStorer
// returns nil if command executed successfully otherwise error
func CreateCar(c car.Car, s db.CarStorer) error {
	return s.StoreCar(c)
}

// Read cars from storage
//
// Pre-cond: given car entity and entity implementing CarStorer interafce
//
// Post-cond: executes command to read cars by given car-pattern from given CarStorer
// returns nil if command executed successfully otherwise error
func ReadCars(pattern car.Car, s db.CarStorer) (std.Linked[car.Car], error) {
	return s.ReadCars(pattern)
}

// Delete cars from storage
//
// Pre-cond: given car entity and entity implementing CarStorer interafce
//
// Post-cond: executes command to delete cars by given car-pattern from given CarStorer
// returns nil if command executed successfully otherwise error
func DeleteCar(pattern car.Car, s db.CarStorer) error {
	return s.DeleteCars(pattern)
}

func SetCar(c car.Car, s db.CarStorer) error {
	return s.SetCar(c)
}

func UnsetCar(c car.Car, s db.CarStorer) error {
	return s.UnsetCar(c)
}
