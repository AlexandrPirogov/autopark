// package kernel collects main functionality for enterprise
package kernel

import (
	"client-service/internal/client"
	"client-service/internal/entity/autopark"
)

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
