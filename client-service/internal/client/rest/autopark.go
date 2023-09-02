package rest

import (
	"client-service/internal/client"
	"client-service/internal/entity/autopark"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

// ReadBrands making request to list avaible brands in autopark-service
//
// Pre-cond: given pattern to list brands that match that pattern
//
// Post-cond: request was executed and result returned.
// If request executes successfully returns list of brands and nil error
// Otherwise returnes nil and error
func (r *Rest) ReadBrands(pattern autopark.Brand) ([]autopark.Brand, error) {

	response, errResp := executeRequest(http.MethodPost, client.ApiGatewayHost+client.BrandListrURL, pattern, r)
	log.Warn().Msgf("got response after request %v, %v", response, errResp)
	if errResp != nil {
		log.Warn().Msgf("%v", errResp)
		return nil, errResp
	}

	defer response.Body.Close()

	brands, unmarshalErr := unmarshalResponse[[]autopark.Brand](response)
	if unmarshalErr != nil {
		log.Warn().Msgf("%v", unmarshalErr)
		return nil, unmarshalErr
	}

	if response.StatusCode == http.StatusOK {
		return brands, nil
	}
	return nil, fmt.Errorf("status code %d", response.StatusCode)
}

// ReadCars making request to list avaible cars in autopark-service
//
// Pre-cond: given pattern to list cars that match that pattern
//
// Post-cond: request was executed and result returned.
// If request executes successfully returns list of cars and nil error
// Otherwise returnes nil and error
func (r *Rest) ReadCars(pattern autopark.Car) ([]autopark.Car, error) {
	response, errResp := executeRequest(http.MethodPost, client.ApiGatewayHost+client.CarListURL, pattern, r)
	if errResp != nil {
		log.Warn().Msgf("%v", errResp)
		return nil, errResp
	}

	defer response.Body.Close()

	cars, unmarshalErr := unmarshalResponse[[]autopark.Car](response)
	if unmarshalErr != nil {
		log.Warn().Msgf("%v", unmarshalErr)
		return nil, unmarshalErr
	}

	if response.StatusCode == http.StatusOK {
		return cars, nil
	}
	return nil, fmt.Errorf("status code %d", response.StatusCode)
}
