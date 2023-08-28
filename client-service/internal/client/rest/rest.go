// package rest holds REST communication with other services
package rest

import (
	"bytes"
	"client-service/internal/client"
	"client-service/internal/entity/autopark"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// New returns pointer to new instance of Rest
//
// Post-cond: returned new pointer to instance of Rest
func New(refreshToken string) *Rest {
	return &Rest{
		client: http.Client{},
		token:  refreshToken,
	}
}

type Rest struct {
	client http.Client
	token  string
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
		log.Println(err)
		return nil, err
	}

	response, errResp := r.executeRequest(http.MethodPost, client.ApiGatewayHost+client.BrandListrURL, payload)
	log.Printf("got response after request %v, %v", response, errResp)
	if errResp != nil {
		log.Println(errResp)
		return nil, errResp
	}

	defer response.Body.Close()

	brands, unmarshalErr := unmarshalResponse[[]autopark.Brand](response)
	if unmarshalErr != nil {
		log.Println(unmarshalErr)
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
	payload, err := json.Marshal(pattern)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	response, errResp := r.executeRequest(http.MethodPost, client.ApiGatewayHost+client.CarListURL, payload)
	if errResp != nil {
		log.Println(errResp)
		return nil, errResp
	}

	defer response.Body.Close()

	cars, unmarshalErr := unmarshalResponse[[]autopark.Car](response)
	if unmarshalErr != nil {
		log.Println(unmarshalErr)
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
	c := http.Cookie{
		Name:  "refresh-token",
		Value: r.token,
	}
	request.AddCookie(&c)
	log.Println("executing request from client ", request.Cookies(), " ", r.token)
	if errReq != nil {
		log.Println(errReq)
		return nil, errReq
	}

	response, errResp := r.client.Do(request)
	if errResp != nil {
		log.Println(errResp)
		return nil, errResp
	}
	return response, nil
}

func unmarshalResponse[T any](response *http.Response) (T, error) {
	var res T
	responseBody, readResponseErr := io.ReadAll(response.Body)
	if readResponseErr != nil {
		log.Println(readResponseErr)
		return res, readResponseErr
	}

	unmarshalErr := json.Unmarshal(responseBody, &res)
	if unmarshalErr != nil {
		log.Println(unmarshalErr)
		return res, unmarshalErr
	}

	return res, nil
}
