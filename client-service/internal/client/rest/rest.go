// package rest holds REST communication with other services
package rest

import (
	"bytes"
	"encoding/json"
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

// executeRequest executes request with given method, url
//
// Pre-cond: given correct http method, url and body
//
// Post-cond: if request was executed successfully return response, nil.
// Otherwise returnes nil, error
func executeRequest[T any](method string, url string, payload T, r *Rest) (*http.Response, error) {
	marshaled, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	reader := bytes.NewReader(marshaled)

	request, errReq := http.NewRequest(method, url, reader)
	c := http.Cookie{
		Name:  "refresh-token",
		Value: r.token,
	}
	request.AddCookie(&c)
	log.Printf("executing request from client to %s with %s", url, marshaled)
	if errReq != nil {
		log.Println(errReq)
		return nil, errReq
	}

	response, errResp := r.client.Do(request)
	if errResp != nil {
		log.Println(errResp)
		return nil, errResp
	}
	log.Printf("got response from %s. Status: %d, err: %v", url, response.StatusCode, errResp)
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
