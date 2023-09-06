// package rest wraps an http.Client for rest requests
package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"manager-service-front/internal/client"
	"net/http"

	"github.com/rs/zerolog/log"
)

type rest struct {
	client http.Client
	token  string
}

// New returns pointer to instance of rest client
//
// Pre-cond: given refresh token value
//
// Post-cond: returned pointer to the new instance of rest client
func New(token string) *rest {
	return &rest{
		http.Client{},
		token,
	}
}

func (r *rest) ListCars() ([]client.Car, error) {
	response, respErr := r.executeRequest(http.MethodPost, client.ApiGatewayHost+client.ListCarsURL, nil)
	if respErr != nil {
		log.Warn().Msgf("%v", respErr)
		return nil, respErr
	}
	defer response.Body.Close()

	res, err := unmarshalResponse[[]client.Car](response)
	if err != nil {
		log.Warn().Msgf("error while unmarshal response %v", err)
		return nil, err
	}
	return res, err
}

func (r *rest) RegisterCar(car []byte) error {
	response, respErr := r.executeRequest(http.MethodPost, client.ApiGatewayHost+client.RegisterCarURL, car)
	if respErr != nil {
		log.Warn().Msgf("%v", respErr)
		return respErr
	}
	defer response.Body.Close()
	return nil
}

func (r *rest) ListBrands() ([]client.Car, error) {
	response, respErr := r.executeRequest(http.MethodPost, client.ApiGatewayHost+client.ListCarsURL, nil)
	if respErr != nil {
		log.Warn().Msgf("%v", respErr)
		return nil, respErr
	}
	defer response.Body.Close()

	res, err := unmarshalResponse[[]client.Car](response)
	if err != nil {
		log.Warn().Msgf("error while unmarshal response %v", err)
		return nil, err
	}
	return res, err
}

func (r *rest) Authenticate(m client.Manager) (*http.Cookie, error) {
	body, _ := json.Marshal(m)
	response, respErr := r.executeRequest(http.MethodPost, client.ApiGatewayHost+client.AuthenticateURL, body)
	if respErr != nil {
		log.Warn().Msgf("%v", respErr)
		return nil, respErr
	}

	c, err := retrieveRefreshToken(response)

	log.Warn().Msgf("token cookie %s %s", c.Name, c.Value)
	defer response.Body.Close()
	return c, err
}

// executeRequest executes request with given method, url
//
// Pre-cond: given correct http method, url and body
//
// Post-cond: if request was executed successfully return response, nil.
// Otherwise returnes nil, error
func (r *rest) executeRequest(method string, url string, body []byte) (*http.Response, error) {
	reader := bytes.NewReader(body)
	log.Warn().Msgf("executing request %s to %s with %s", method, url, body)
	request, reqErr := http.NewRequest(http.MethodPost, url, reader)
	c := http.Cookie{
		Name:  client.RerfeshTokenCookieField,
		Value: r.token,
	}
	request.AddCookie(&c)
	if reqErr != nil {
		log.Warn().Msgf("%v", reqErr)
		return nil, reqErr
	}

	response, errResp := r.client.Do(request)
	if errResp != nil {
		log.Warn().Msgf("%v", errResp)
		return nil, errResp
	}
	return response, nil
}

func retrieveRefreshToken(r *http.Response) (*http.Cookie, error) {
	for _, c := range r.Cookies() {
		if c.Name == "refresh-token" {
			return c, nil
		}
	}
	return nil, fmt.Errorf("cookie not found")
}

func unmarshalResponse[T any](response *http.Response) (T, error) {
	var res T
	responseBody, readResponseErr := io.ReadAll(response.Body)
	log.Warn().Msgf("got response body %s with status %d", responseBody, response.StatusCode)
	log.Warn().Msgf("response cookie %v", response.Cookies())
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
