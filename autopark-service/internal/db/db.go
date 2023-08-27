package db

import (
	"autopark-service/internal/autopark/car"
	"autopark-service/internal/db/postgres"
	"autopark-service/internal/std"
)

type BrandStorer interface {
	StoreBrand(b car.Brand) error
	ReadBrands(pattern car.Brand) (std.Linked[car.Brand], error)
}

type CarStorer interface {
	StoreCar(c car.Car) error
	DeleteCars(pattern car.Car) error
	ReadCars(pattern car.Car) (std.Linked[car.Car], error)
}

type Storer interface {
	BrandStorer
	CarStorer
}

func GetConnInstance() Storer {
	return postgres.GetInstance()
}
