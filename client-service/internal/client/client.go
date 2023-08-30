// package client holds interface and URLs to communicate with other services
package client

import (
	"client-service/internal/entity"
	"client-service/internal/entity/autopark"
)

// ApiGateway URL
const ApiGatewayHost = "http://api-gateway-nginx"

// Lists brandss
const BrandListrURL = "/autopark/brand/list"

// Register new brand
const BrandRegisterURL = "/autopark/brand/register"

// List cars
const CarListURL = "/autopark/car/set/list"

// Register new acrs
const CarRegisterURL = "/autopark/car/register"

// Delete car
const CarDeleteURL = "/autopark/car/remove"

// Booking URLs
// create booking
const BookingCreateURL = "/booking/create"

// approve booking
const BookingApproveURL = "/booking/approve"

// cancel booking
const BookingCancelURL = "/booking/cancel"

// choose car for booking
const BookingChooseURL = "/booking/choose"

// Finish booking
const BookingFinishURL = "/booking/finish"

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

type ClientBooking interface {
	// BookingApprove send request to service booking to approve given booking
	//
	// Pre-cond: given entity.Booking to approve
	//
	// Post-cond: request was executed. If successfull returns nil otherwise error
	BookingApprove(entity.Booking) (entity.Booking, error)

	// BookingCreate send request to service booking to create new booking for given entity.ClientCreds
	//
	// Pre-cond: given entity.ClientCreds to create new booking for
	//
	// Post-cond: request was executed. If successfull returns booking with set id and nil; otherwise error
	BookingCreate(entity.ClientCreds) (entity.Booking, error)

	// BookingCancel send request to service booking to cancel given booking
	//
	// Pre-cond: given entity.Booking to cancel
	//
	// Post-cond: request was executed. If successfull returns nil otherwise error
	BookingCancel(entity.Booking) error

	// BookingChoose send request to service booking to choose car for given booking
	//
	// Pre-cond: given entity.Booking to choose car with set Car ID
	//
	// Post-cond: request was executed. If successfull returns nil otherwise error
	BookingChoose(entity.Booking) error

	// BookingFinish send request to service booking to finish given booking
	//
	// Pre-cond: given entity.Booking to finish
	//
	// Post-cond: request was executed. If successfull returns nil otherwise error
	BookingFinish(entity.Booking) error
}
