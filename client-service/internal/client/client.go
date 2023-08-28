// package client holds interface and URLs to communicate with other services
package client

import "client-service/internal/entity/autopark"

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

const ClientRegisterURL = "/auth/register/client"

// ClientAutopark interface is the main interface to communicate with autopark service
type ClientAutopark interface {

	// ReadBrands making request to list avaible brands in autopark-service
	//
	// Pre-cond: given pattern to list brands that match that pattern
	//
	// Post-cond: request was executed and result returned.
	// If request executes successfully returns list of brands and nil error
	// Otherwise returnes nil and error
	ReadBrands(pattern autopark.Brand) ([]autopark.Brand, error)

	// ReadCars making request to list avaible cars in autopark-service
	//
	// Pre-cond: given pattern to list cars that match that pattern
	//
	// Post-cond: request was executed and result returned.
	// If request executes successfully returns list of cars and nil error
	// Otherwise returnes nil and error
	ReadCars(pattern autopark.Car) ([]autopark.Car, error)
}
