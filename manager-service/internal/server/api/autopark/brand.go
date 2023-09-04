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

// BrandList is API for listing brands in autopark-service
func BrandList(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Warn().Msgf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pattern, unmarshalErr := unmarshal[autopark.Brand](body)
	if unmarshalErr != nil {
		log.Printf("err while unmarshal reequest body %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, requestErr := kernel.ReadBrands(pattern, rest.New())
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

// BrandRegister is API for register new brand in autopark-service
func BrandRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("err while reading reequest body %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, unmarshalErr := unmarshal[autopark.Brand](body)
	if unmarshalErr != nil {
		log.Printf("err while unmarshal reequest body %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respErr := kernel.StoreBrand(b, rest.New())
	if respErr == nil {
		w.WriteHeader(http.StatusCreated)
		return
	}

	log.Printf("err while making http request %v", err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(respErr.Error()))
}
