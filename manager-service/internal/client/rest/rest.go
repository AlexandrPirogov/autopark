// package rest holds REST communication with other services
package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"manager-service/internal/client"
	"manager-service/internal/entity/autopark"
	"net/http"

	"github.com/rs/zerolog/log"
)

// New returns pointer to new instance of Rest
//
// Post-cond: returned new pointer to instance of Rest
func New() *Rest {
	return &Rest{
		client: http.Client{},
	}
}

type Rest struct {
	client http.Client
}

// StoreBrand making request to store given brand in autopark-service
//
// Pre-cond: given brand to store
//
// Post-cond: request was executed and result returned.
// If brand was stored successfully then returns nil, otherwise returns error
func (r *Rest) StoreBrand(b autopark.Brand) error {
	payload, err := json.Marshal(b)
	if err != nil {
		log.Warn().Msgf("%v", err)
		return err
	}

	response, errResp := r.executeRequest(http.MethodPost, client.ApiGatewayHost+client.BrandRegisterURL, payload)
	if errResp != nil {
		log.Warn().Msgf("%v", errResp)
		return errResp
	}

	defer response.Body.Close()
	if response.StatusCode == http.StatusCreated {
		return nil
	}
	return fmt.Errorf("status code %d", response.StatusCode)
}

// ReadBrands making request to list avaible brands in autopark-service
//
// Pre-cond: given pattern to list brands that match that pattern
//
// Post-cond: request was executed and result returned.
// If request executes successfully returns list of brands and nil error
// Otherwise returnes nil and error
func (r *Rest) ReadBrands(pattern autopark.Brand) ([]autopark.Brand, error) {
	payload, err := json.Marshal(pattern)
	if err != nil {
		log.Warn().Msgf("%v", err)
		return nil, err
	}

	response, errResp := r.executeRequest(http.MethodPost, client.ApiGatewayHost+client.BrandListrURL, payload)
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

// StoreCar making request to store given car in autopark-service
//
// Pre-cond: given car to store
//
// Post-cond: request was executed and result returned.
// If car was stored successfully then returns nil, otherwise returns error
func (r *Rest) StoreCar(c autopark.Car) error {
	payload, err := json.Marshal(c)
	if err != nil {
		log.Warn().Msgf("%v", err)
		return err
	}

	response, errResp := r.executeRequest(http.MethodPost, client.ApiGatewayHost+client.CarRegisterURL, payload)
	if errResp != nil {
		log.Warn().Msgf("%v", errResp)
		return errResp
	}

	defer response.Body.Close()
	if response.StatusCode == http.StatusCreated {
		return nil
	}
	return fmt.Errorf("status code %d", response.StatusCode)
}

// DeleteCars making request to delete cars from autopark-service with given pattern
//
// Pre-cond: given car pattern
//
// Post-cond: request was executed.
// If request was executed successfully then returns status nil error, otherwise returns error
func (r *Rest) DeleteCars(pattern autopark.Car) error {
	payload, err := json.Marshal(pattern)
	if err != nil {
		log.Warn().Msgf("%v", err)
		return err
	}

	response, errResp := r.executeRequest(http.MethodPost, client.ApiGatewayHost+client.CarDeleteURL, payload)
	if errResp != nil {
		log.Warn().Msgf("%v", errResp)
		return errResp
	}

	defer response.Body.Close()
	if response.StatusCode == http.StatusNoContent {
		return nil
	}
	return fmt.Errorf("status code %d", response.StatusCode)
}

// ReadCars making request to list avaible cars in autopark-service
//
// Pre-cond: given pattern to list cars that match that pattern
//
// Post-cond: request was executed and result returned.
// If request executes successfully returns list of cars and nil error
// Otherwise returnes nil and error
func (r *Rest) ReadCars(pattern autopark.Car) ([]autopark.Car, error) {
	payload, err := json.Marshal(pattern)
	if err != nil {
		log.Warn().Msgf("%v", err)
		return nil, err
	}

	response, errResp := r.executeRequest(http.MethodPost, client.ApiGatewayHost+client.CarListURL, payload)
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

// executeRequest executes request with given method, url
//
// Pre-cond: given correct http method, url and body
//
// Post-cond: if request was executed successfully return response, nil.
// Otherwise returnes nil, error
func (r *Rest) executeRequest(method string, url string, body []byte) (*http.Response, error) {
	reader := bytes.NewReader(body)
	request, errReq := http.NewRequest(http.MethodPost, url, reader)
	if errReq != nil {
		log.Warn().Msgf("%v", errReq)
		return nil, errReq
	}

	response, errResp := r.client.Do(request)
	if errResp != nil {
		log.Warn().Msgf("%v", errResp)
		return nil, errResp
	}
	return response, nil
}

func unmarshalResponse[T any](response *http.Response) (T, error) {
	var res T
	responseBody, readResponseErr := io.ReadAll(response.Body)
	if readResponseErr != nil {
		log.Warn().Msgf("%v", readResponseErr)
		return res, readResponseErr
	}

	unmarshalErr := json.Unmarshal(responseBody, &res)
	if unmarshalErr != nil {
		log.Warn().Msgf("%v", unmarshalErr)
		return res, unmarshalErr
	}

	return res, nil
}
