// package rest wraps an http.Client for rest requests
package rest

import (
	"bytes"
	"encoding/json"
	"enterprise-front/internal/client"
	"io"
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

// RegisterManager making request to register manager in auth-service
//
// Pre-cond: given client.Manager instance to register
//
// Post-cond: request was executed and result returned.
// If request executes successfully returns Manager that was registeres and nil error
// Otherwise returnes nil and error
func (r *rest) RegisterManager(m client.Manager) (client.Manager, error) {
	var res client.Manager
	body, marshalErr := json.Marshal(m)
	if marshalErr != nil {
		log.Warn().Msgf("%v", marshalErr)
		return res, marshalErr
	}

	response, respErr := r.executeRequest(http.MethodPost, client.ApiGatewayHost+client.RegisterEnterpriseManagerURL, body)
	if respErr != nil {
		log.Warn().Msgf("%v", respErr)
		return res, respErr
	}
	defer response.Body.Close()

	res, unmarshalErr := unmarshalResponse[client.Manager](response)
	if unmarshalErr != nil {
		log.Warn().Msgf("%v", unmarshalErr)
		return res, unmarshalErr
	}

	return res, nil
}

// RegisterManager making request to register manager in auth-service
//
// Pre-cond: given client.Manager instance to register
//
// Post-cond: request was executed and result returned.
// If request executes successfully returns Manager that was registeres and nil error
// Otherwise returnes nil and error
func (r *rest) RegisterEnterprise(m client.Enterprise) error {
	body, marshalErr := json.Marshal(m)
	if marshalErr != nil {
		log.Warn().Msgf("%v", marshalErr)
		return marshalErr
	}

	response, respErr := r.executeRequest(http.MethodPost, client.ApiGatewayHost+client.RegisterEnterpriseURL, body)
	if respErr != nil {
		log.Warn().Msgf("%v", respErr)
		return respErr
	}
	defer response.Body.Close()

	return nil
}

func (r *rest) ListEnterprises() ([]client.Enterprise, error) {
	response, respErr := r.executeRequest(http.MethodPost, client.ApiGatewayHost+client.ListEnterprisesURL, []byte{})
	if respErr != nil {
		log.Warn().Msgf("%v", respErr)
		return []client.Enterprise{}, respErr
	}
	defer response.Body.Close()
	res, err := unmarshalResponse[[]client.Enterprise](response)
	if err != nil {
		return []client.Enterprise{}, err
	}
	return res, nil
}

func (r *rest) EnterpriseByTitle(title string) (client.Enterprise, error) {
	response, respErr := r.executeRequest(http.MethodPost, client.ApiGatewayHost+client.EnterpriseByTitle+title, nil)
	if respErr != nil {
		log.Warn().Msgf("%v", respErr)
		return client.Enterprise{}, respErr
	}
	defer response.Body.Close()
	res, err := unmarshalResponse[client.Enterprise](response)
	if err != nil {
		return client.Enterprise{}, err
	}
	return res, nil
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

func unmarshalResponse[T any](response *http.Response) (T, error) {
	var res T
	responseBody, readResponseErr := io.ReadAll(response.Body)
	log.Warn().Msgf("got response body %s with status %d", responseBody, response.StatusCode)
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
