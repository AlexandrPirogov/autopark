package api

import (
	"encoding/json"
	"enterprise-service/internal/client"
	"enterprise-service/internal/client/rest"
	"enterprise-service/internal/db"
	"enterprise-service/internal/enterprise"
	"enterprise-service/internal/kernel"
	"enterprise-service/internal/std"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	marhsalErr := json.Unmarshal(body, &e)
	if marhsalErr != nil {
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
	var e enterprise.Enterprise
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	marhsalErr := json.Unmarshal(body, &e)
	if marhsalErr != nil {
		log.Println(marhsalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	list, storeErr := kernel.Read(e, db.GetConnInstance())
	if storeErr != nil {
		log.Println(storeErr)
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
	idstr := chi.URLParam(r, "id")
	id, convErr := strconv.Atoi(idstr)
	if convErr != nil {
		log.Println(convErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, readErr := kernel.ReadByID(id, db.GetConnInstance())
	if readErr != nil {
		log.Println(readErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	responseBody, marshalErr := json.Marshal(res)
	if marshalErr != nil {
		log.Println(marshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func RegisterManager(w http.ResponseWriter, r *http.Request) {
	idstr := chi.URLParam(r, "id")
	id, convErr := strconv.Atoi(idstr)
	if convErr != nil {
		log.Println(convErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var m client.Manager
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	marhsalErr := json.Unmarshal(body, &m)
	if marhsalErr != nil {
		log.Println(marhsalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, _ := retrieveRefreshToken(r)
	res, err := kernel.RegisterManager(m, rest.New(token))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res.EnterpriseID = id
	assignErr := kernel.AssignManager(res, db.GetConnInstance())
	if assignErr != nil {
		log.Println(assignErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("register response %v", res)
	responseBody, marshalErr := json.Marshal(res)
	if marshalErr != nil {
		log.Println(marshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)
}

func retrieveRefreshToken(r *http.Request) (string, error) {
	var cookie *http.Cookie
	cookie, err := r.Cookie(client.RerfeshTokenCookieField)

	if cookie == nil || err != nil {
		log.Println(r.Header[client.RerfeshTokenCookieField])
		return "", fmt.Errorf("error while reading token %v", err)
	}

	tokenVal := cookie.Value[strings.Index(cookie.Value, "=")+1:]

	return tokenVal, err
}
