// package kernel collects main functionality for enterprise
package kernel

import (
	"manager-service/internal/client"
	"manager-service/internal/entity/autopark"
)

// StoreBrand makes given client to  to store given brand in autopark-service
//
// Pre-cond: given brand to store and client to communicate with autopark-service
//
// Post-cond: request was executed and result returned.
// If brand was stored successfully then returns nil, otherwise returns error
func StoreBrand(b autopark.Brand, c client.ClientAutopark) error {
	return c.StoreBrand(b)
}

// ReadBrands making given client to request to list avaible brands in autopark-service
//
// Pre-cond: given pattern to list brands that match that pattern nd client to communicate with autopark-service
//
// Post-cond: request was executed and result returned.
// If request executes successfully returns list of brands and nil error
// Otherwise returnes nil and error
func ReadBrands(pattern autopark.Brand, c client.ClientAutopark) ([]autopark.Brand, error) {
	return c.ReadBrands(pattern)
}

// StoreCar making given client to request to store given car in autopark-service
//
// Pre-cond: given car to store and client to communicate with autopark-service
//
// Post-cond: request was executed and result returned.
// If car was stored successfully then returns nil, otherwise returns error
func StoreCar(car autopark.Car, c client.ClientAutopark) error {
	return c.StoreCar(car)
}

// DeleteCars making client to request to delete cars from autopark-service with given pattern
//
// Pre-cond: given car pattern and client to communicate with autopark-service
//
// Post-cond: request was executed.
// If request was executed successfully then returns status nil error, otherwise returns error
func DeleteCars(pattern autopark.Car, c client.ClientAutopark) error {
	return c.DeleteCars(pattern)
}

// ReadCars making client to request to list avaible cars in autopark-service
//
// Pre-cond: given pattern to list cars that match that pattern and client to communicate with autopark-service
//
// Post-cond: request was executed and result returned.
// If request executes successfully returns list of cars and nil error
// Otherwise returnes nil and error
func ReadCars(pattern autopark.Car, c client.ClientAutopark) ([]autopark.Car, error) {
	return c.ReadCars(pattern)
}
