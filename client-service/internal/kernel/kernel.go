// package kernel collects main functionality for enterprise
package kernel

import (
	"client-service/internal/client"
	"client-service/internal/entity"
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

// CreateBooknig making a request to create new book in booking-service
//
// Pre-cond: given entity.ClientCreds to create new booking for and implemtation of client.ClientBooking
//
// Post-cond: client.ClientBooking executes request and returns result of request
func CreateBooking(e entity.ClientCreds, c client.ClientBooking) (entity.Booking, error) {
	return c.BookingCreate(e)
}

// ChooseCarBooking making a request to choose car for given booking in booking-service
//
// Pre-cond: given entity.Booking to choose car for and implemtation of client.ClientBooking
//
// Post-cond: client.ClientBooking executes request and returns result of request
func ChooseCarBooking(b entity.Booking, c client.ClientBooking) error {
	return c.BookingChoose(b)
}

// ApproveBooking making a request to approve given book in booking-service
//
// Pre-cond: given entity.Booking to approve and implemtation of client.ClientBooking
//
// Post-cond: client.ClientBooking executes request and returns result of request
func ApproveBooking(b entity.Booking, c client.ClientBooking) (entity.Booking, error) {
	return c.BookingApprove(b)
}

// CancelBooking making a request to cancel given book in booking-service
//
// Pre-cond: given entity.Booking to cancel and implemtation of client.ClientBooking
//
// Post-cond: client.ClientBooking executes request and returns result of request
func CancelBooking(b entity.Booking, c client.ClientBooking) error {
	return c.BookingCancel(b)
}

// FinishBooking making a request to finish given book in booking-service
//
// Pre-cond: given entity.Booking to finish implemtation of client.ClientBooking
//
// Post-cond: client.ClientBooking executes request and returns result of request
func FinishBooking(b entity.Booking, c client.ClientBooking) error {
	return c.BookingFinish(b)
}
