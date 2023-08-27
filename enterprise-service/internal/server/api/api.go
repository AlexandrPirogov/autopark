package api

import (
	"encoding/json"
	"enterprise-service/internal/client"
	"enterprise-service/internal/client/rest"
	"enterprise-service/internal/db"
	"enterprise-service/internal/enterprise"
	"enterprise-service/internal/kernel"
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

	responseBody := []byte{}
	for {
		item, ok := list.PopFront()
		if !ok {
			break
		}

		marshaled, err := json.Marshal(item)
		if err != nil {
			continue
		}

		responseBody = append(responseBody, marshaled...)
	}

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

// Update updates existing enterprise entity in system
//
// Pre-cond: given enterprise pattern to update and new params for them
//
// Post-cond: all enterprises that matches the given pattern was update
// with given new params
func UpdateEnterprises(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func retrieveRefreshToken(r *http.Request) (string, error) {
	var cookie *http.Cookie
	cookie, err := r.Cookie("refresh-token")

	if cookie == nil || err != nil {
		log.Println(r.Header[client.RerfeshTokenCookieField])
		return "", fmt.Errorf("error while reading token %v", err)
	}

	tokenVal := cookie.Value[strings.Index(cookie.Value, "=")+1:]

	return tokenVal, err
}
