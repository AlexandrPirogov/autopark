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
		log.Warn().Msgf("err while unmarshal reequest body %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, _ := api.RetrieveRefreshToken(r)
	log.Warn().Msgf("executing request %s with token %s", body, token)
	response, requestErr := kernel.ReadBrands(pattern, rest.New(token))
	log.Warn().Msgf("got response %v, %v", response, err)
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
