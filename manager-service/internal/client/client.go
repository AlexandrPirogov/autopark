package client

import (
	"manager-service/internal/entity/autopark"
	"manager-service/internal/std"
)

type ClientAutopark interface {
	StoreBrand(b autopark.Brand) error
	ReadBrands(pattern autopark.Brand) (std.Linked[autopark.Brand], error)

	StoreCar(c autopark.Car) error
	DeleteCars(pattern autopark.Car) error
	ReadCars(pattern autopark.Car) (std.Linked[autopark.Car], error)
}
