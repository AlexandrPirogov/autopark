package autopark

import (
	"encoding/json"
	"io"
	"manager-service/internal/client/rest"
	"manager-service/internal/entity/autopark"
	"manager-service/internal/kernel"
	"net/http"

	"github.com/rs/zerolog/log"
)

// CarList is API for listing brands in autopark-service
func CarList(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Warn().Msgf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pattern, unmarshalErr := unmarshal[autopark.Car](body)
	if unmarshalErr != nil {
		log.Printf("err while unmarshal reequest body %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, requestErr := kernel.ReadCars(pattern, rest.New())
	if requestErr != nil {
		log.Warn().Msgf("%v", requestErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, marshalErr := json.Marshal(response)
	if marshalErr != nil {
		log.Warn().Msgf("%v", marshalErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

// CarRegister is API to register new car in autopark-service
func CarRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("err while reading reequest body %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c, unmarshalErr := unmarshal[autopark.Car](body)
	if unmarshalErr != nil {
		log.Printf("err while unmarshal reequest body %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respErr := kernel.StoreCar(c, rest.New())
	if respErr == nil {
		w.WriteHeader(http.StatusCreated)
		return
	}

	log.Printf("err while making http request %v", err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(respErr.Error()))
}

// CarDelete is API to delete car in autopark-service
func CarDelete(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("err while reading reequest body %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c, unmarshalErr := unmarshal[autopark.Car](body)
	if unmarshalErr != nil {
		log.Printf("err while unmarshal reequest body %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respErr := kernel.DeleteCars(c, rest.New())
	if respErr == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	log.Printf("err while making http request %v", err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(respErr.Error()))
}
