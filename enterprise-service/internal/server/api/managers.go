package api

import (
	"encoding/json"
	"enterprise-service/internal/client"
	"enterprise-service/internal/client/rest"
	"enterprise-service/internal/db"
	"enterprise-service/internal/kernel"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

// RegisterManagers regsiter new manager in auth-service
func RegisterManager(w http.ResponseWriter, r *http.Request) {

	var m client.Manager
	body, err := io.ReadAll(r.Body)
	log.Warn().Msgf("register manager got %s", body)
	if err != nil {
		log.Warn().Msgf("err while reading %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	marhsalErr := json.Unmarshal(body, &m)
	if marhsalErr != nil {
		log.Warn().Msgf("err while unmarshal %v", marhsalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, _ := retrieveRefreshToken(r)
	res, err := kernel.RegisterManager(m, rest.New(token))
	if err != nil {
		log.Warn().Msgf("err while registering %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	assignErr := kernel.AssignManager(m, db.GetConnInstance())
	if assignErr != nil {
		log.Warn().Msgf("assign error %v", assignErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	m.Id = res.Id
	log.Printf("register response %v", m)
	responseBody, marshalErr := json.Marshal(m)
	if marshalErr != nil {
		log.Warn().Msgf("%v", marshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)
}
