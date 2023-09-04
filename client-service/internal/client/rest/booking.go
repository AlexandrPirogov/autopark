package rest

import (
	"client-service/internal/client"
	"client-service/internal/entity"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

// BookingApprove send request to service booking to approve given booking
//
// Pre-cond: given entity.Booking to approve
//
// Post-cond: request was executed. If successfull returns nil otherwise error
func (r *Rest) BookingApprove(b entity.Booking) (entity.Booking, error) {
	var createdBooking entity.Booking
	response, errResp := executeRequest(http.MethodPost, client.ApiGatewayHost+client.BookingApproveURL, b, r)
	log.Warn().Msgf("got response after request %v, %v", response, errResp)
	if errResp != nil {
		log.Warn().Msgf("%v", errResp)
		return createdBooking, errResp
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		return createdBooking, nil
	}
	return createdBooking, fmt.Errorf("status code %d", response.StatusCode)
}

// BookingCreate send request to service booking to create new booking for given entity.ClientCreds
//
// Pre-cond: given entity.ClientCreds to create new booking for
//
// Post-cond: request was executed. If successfull returns booking with set id and nil; otherwise error
func (r *Rest) BookingCreate(e entity.ClientCreds) (entity.Booking, error) {
	var createdBooking entity.Booking

	response, errResp := executeRequest(http.MethodPost, client.ApiGatewayHost+client.BookingCreateURL, e, r)
	log.Warn().Msgf("got response after request %v, %v", response, errResp)
	if errResp != nil {
		log.Warn().Msgf("%v", errResp)
		return createdBooking, errResp
	}

	defer response.Body.Close()

	createdBooking, unmarshalErr := unmarshalResponse[entity.Booking](response)
	if unmarshalErr != nil {
		log.Warn().Msgf("%v", unmarshalErr)
		return createdBooking, unmarshalErr
	}

	if response.StatusCode == http.StatusOK {
		return createdBooking, nil
	}
	return createdBooking, fmt.Errorf("status code %d", response.StatusCode)
}

// BookingCancel send request to service booking to cancel given booking
//
// Pre-cond: given entity.Booking to cancel
//
// Post-cond: request was executed. If successfull returns nil otherwise error
func (r *Rest) BookingCancel(b entity.Booking) error {

	response, errResp := executeRequest(http.MethodPost, client.ApiGatewayHost+client.BookingCancelURL, b, r)
	log.Warn().Msgf("got response after request %v, %v", response, errResp)
	if errResp != nil {
		log.Warn().Msgf("%v", errResp)
		return errResp
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		return nil
	}
	return fmt.Errorf("status code %d", response.StatusCode)
}

// BookingChoose send request to service booking to choose car for given booking
//
// Pre-cond: given entity.Booking to choose car with set Car ID
//
// Post-cond: request was executed. If successfull returns nil otherwise error
func (r *Rest) BookingChoose(b entity.Booking) error {

	response, errResp := executeRequest(http.MethodPost, client.ApiGatewayHost+client.BookingChooseURL, b, r)
	log.Warn().Msgf("got response after request %v, %v", response, errResp)
	if errResp != nil {
		log.Warn().Msgf("%v", errResp)
		return errResp
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		return nil
	}
	return fmt.Errorf("status code %d", response.StatusCode)
}

// BookingFinish send request to service booking to finish given booking
//
// Pre-cond: given entity.Booking to finish
//
// Post-cond: request was executed. If successfull returns nil otherwise error
func (r *Rest) BookingFinish(b entity.Booking) error {
	response, errResp := executeRequest(http.MethodPost, client.ApiGatewayHost+client.BookingFinishURL, b, r)
	log.Warn().Msgf("got response after request %v, %v", response, errResp)
	if errResp != nil {
		log.Warn().Msgf("%v", errResp)
		return errResp
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		return nil
	}
	return fmt.Errorf("status code %d", response.StatusCode)
}
