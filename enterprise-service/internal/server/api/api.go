package api

import (
	"encoding/json"
	"enterprise-service/internal/client"
	"enterprise-service/internal/db"
	"enterprise-service/internal/enterprise"
	"enterprise-service/internal/kernel"
	"enterprise-service/internal/std"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/go-chi/chi/v5"
)

// RegisterEnterprise creates new enterprise entity in system
//
// Pre-cond: given unique enterprise entity
//
// Post-cond: new enterprise entity for created for system
func RegisterEnterprise(w http.ResponseWriter, r *http.Request) {
	var e enterprise.Enterprise
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Warn().Msgf("err while reading response %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	marhsalErr := json.Unmarshal(body, &e)
	if marhsalErr != nil {
		log.Warn().Msgf("err while unmarshal %v", marhsalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storeErr := kernel.Store(e, db.GetConnInstance())
	if storeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Delete removes given enterprise entity from system
//
// Pre-cond: given existing enterprise to remove
//
// Post-cond: enterprise was removed from system
func DeleteEnterprise(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// Read read all existing enterprises in system by given pattern
//
// Pre-cond: given pattern for enterprises
//
// Post-cond: returns list of matched enterprises
func ReadEnerprises(w http.ResponseWriter, r *http.Request) {
	list, storeErr := kernel.Read(db.GetConnInstance())
	if storeErr != nil {
		log.Warn().Msgf("%v", storeErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	toMarshal := std.AsSlice(list.Iterator())

	responseBody, _ := json.Marshal(toMarshal)

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

// Read read all existing enterprises in system by given pattern
//
// Pre-cond: given pattern for enterprises
//
// Post-cond: returns list of matched enterprises
func ReadEnerprise(w http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r, "title")

	res, readErr := kernel.ReadByTitle(title, db.GetConnInstance())
	if readErr != nil {
		log.Warn().Msgf("%v", readErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	responseBody, marshalErr := json.Marshal(res)
	if marshalErr != nil {
		log.Warn().Msgf("%v", marshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func retrieveRefreshToken(r *http.Request) (string, error) {
	var cookie *http.Cookie
	cookie, err := r.Cookie(client.RerfeshTokenCookieField)

	if cookie == nil || err != nil {
		log.Warn().Msgf("%v", r.Header[client.RerfeshTokenCookieField])
		return "", fmt.Errorf("error while reading token %v", err)
	}

	tokenVal := cookie.Value[strings.Index(cookie.Value, "=")+1:]

	return tokenVal, err
}
