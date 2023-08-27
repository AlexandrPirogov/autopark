package rest

import (
	"bytes"
	"encoding/json"
	"enterprise-service/internal/client"
	"io"
	"log"
	"net/http"
)

type RestClient struct {
	client http.Client
	token  string
}

func New(token string) *RestClient {
	return &RestClient{
		http.Client{},
		token,
	}
}

func (r *RestClient) RegisterManager(m client.Manager) (client.Manager, error) {
	var res client.Manager
	body, marshalErr := json.Marshal(m)
	if marshalErr != nil {
		log.Println(marshalErr)
		return res, marshalErr
	}

	reader := bytes.NewReader(body)
	log.Printf("sending %s", body)
	request, reqErr := http.NewRequest(http.MethodPost, client.ApiGatewayHost+client.RegisterManagerURL, reader)
	request.Header.Add("refresh-token", r.token)
	if reqErr != nil {
		log.Println(reqErr)
		return res, reqErr
	}

	response, respErr := r.client.Do(request)
	if respErr != nil {
		log.Println(respErr)
		return res, respErr
	}
	log.Printf("got %s", response.Body)

	defer response.Body.Close()
	responseBody, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		log.Println(readErr)
		return res, readErr
	}

	unmarshalErr := json.Unmarshal(responseBody, &res)
	if unmarshalErr != nil {
		log.Println(unmarshalErr)
		return res, unmarshalErr
	}

	return res, nil
}
