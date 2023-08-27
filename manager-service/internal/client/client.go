// package client holds interface and URLs to communicate with other services
package client

import (
	"manager-service/internal/entity/autopark"
)

// ApiGateway URL
const ApiGatewayHost = "http://api-gateway-nginx"

// Lists brandss
const BrandListrURL = "/autopark/brand/list"

// Register new brand
const BrandRegisterURL = "/autopark/brand/register"

// List cars
const CarListURL = "/autopark/car/list"

// Register new acrs
const CarRegisterURL = "/autopark/car/register"

// Delete car
const CarDeleteURL = "/autopark/car/remove"

// ClientAutopark interface is the main interface to communicate with autopark service
type ClientAutopark interface {
	// StoreBrand making request to store given brand in autopark-service
	//
	// Pre-cond: given brand to store
	//
	// Post-cond: request was executed and result returned.
	// If brand was stored successfully then returns nil, otherwise returns error
	StoreBrand(b autopark.Brand) error

	// ReadBrands making request to list avaible brands in autopark-service
	//
	// Pre-cond: given pattern to list brands that match that pattern
	//
	// Post-cond: request was executed and result returned.
	// If request executes successfully returns list of brands and nil error
	// Otherwise returnes nil and error
	ReadBrands(pattern autopark.Brand) ([]autopark.Brand, error)

	// StoreCar making request to store given car in autopark-service
	//
	// Pre-cond: given car to store
	//
	// Post-cond: request was executed and result returned.
	// If car was stored successfully then returns nil, otherwise returns error
	StoreCar(c autopark.Car) error

	// DeleteCars making request to delete cars from autopark-service with given pattern
	//
	// Pre-cond: given car pattern
	//
	// Post-cond: request was executed.
	// If request was executed successfully then returns status nil error, otherwise returns error
	DeleteCars(pattern autopark.Car) error

	// ReadCars making request to list avaible cars in autopark-service
	//
	// Pre-cond: given pattern to list cars that match that pattern
	//
	// Post-cond: request was executed and result returned.
	// If request executes successfully returns list of cars and nil error
	// Otherwise returnes nil and error
	ReadCars(pattern autopark.Car) ([]autopark.Car, error)
}
