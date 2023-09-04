package autopark

import (
	"client-service/internal/client/rest"
	"client-service/internal/entity/autopark"
	"client-service/internal/kernel"
	"client-service/internal/server/api"
	"encoding/json"
	"io"
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
		log.Warn().Msgf("err while unmarshal reequest body %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, _ := api.RetrieveRefreshToken(r)
	response, requestErr := kernel.ReadCars(pattern, rest.New(token))
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
