// package kernel collects main functionality for enterprise
package kernel

import (
	"manager-service/internal/client"
	"manager-service/internal/entity/autopark"
	"manager-service/internal/std"
)

func StoreBrand(b autopark.Brand, c client.ClientAutopark) error {
	return c.StoreBrand(b)
}

func ReadBrands(pattern autopark.Brand, c client.ClientAutopark) (std.Linked[autopark.Brand], error) {
	return c.ReadBrands(pattern)
}

func StoreCar(car autopark.Car, c client.ClientAutopark) error {
	return c.StoreCar(car)
}

func DeleteCars(pattern autopark.Car, c client.ClientAutopark) error {
	return c.DeleteCars(pattern)
}

func ReadCars(pattern autopark.Car, c client.ClientAutopark) (std.Linked[autopark.Car], error) {
	return c.ReadCars(pattern)
}
