package rest

import (
	"manager-service/internal/entity/autopark"
	"manager-service/internal/std"
	"net/http"
)

func New() Rest {
	return Rest{
		client: http.Client{},
	}
}

type Rest struct {
	client http.Client
}

func (r *Rest) StoreBrand(b autopark.Brand) error {
	return nil
}

func (r *Rest) ReadBrands(pattern autopark.Brand) (std.Linked[autopark.Brand], error) {
	return nil, nil
}

func (r *Rest) StoreCar(c autopark.Car) error {
	return nil
}

func (r *Rest) DeleteCars(pattern autopark.Car) error {
	return nil
}

func (r *Rest) ReadCars(pattern autopark.Car) (std.Linked[autopark.Car], error) {
	return nil, nil
}
