package booking

import (
	"client-service/internal/client/rest"
	"client-service/internal/entity"
	"client-service/internal/kernel"
	"client-service/internal/server/api"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Create is API for creating new  booking that marshaled from request body for user
func Create(w http.ResponseWriter, r *http.Request) {
	c, err := readRequestBody[entity.ClientCreds](r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, _ := api.RetrieveRefreshToken(r)
	created, bookingErr := kernel.CreateBooking(c, rest.New(token))

	if bookingErr != nil {
		log.Println(bookingErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	responseBody, marshalErr := json.Marshal(created)
	if marshalErr != nil {
		log.Println(marshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)

}

// Choose is API for choosing car for given booking that marshaled from request body
func Choose(w http.ResponseWriter, r *http.Request) {
	u, err := readRequestBody[entity.Booking](r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, _ := api.RetrieveRefreshToken(r)
	bookingErr := kernel.ChooseCarBooking(u, rest.New(token))

	if bookingErr != nil {
		log.Println(bookingErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Cancel is API for canceling booking that marshaled from request body
func Cancel(w http.ResponseWriter, r *http.Request) {
	u, err := readRequestBody[entity.Booking](r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, _ := api.RetrieveRefreshToken(r)
	bookingErr := kernel.CancelBooking(u, rest.New(token))

	if bookingErr != nil {
		log.Println(bookingErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Approve is API for approving booking that marshaled from request body
func Approve(w http.ResponseWriter, r *http.Request) {

	u, err := readRequestBody[entity.Booking](r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, _ := api.RetrieveRefreshToken(r)
	_, bookingErr := kernel.ApproveBooking(u, rest.New(token))

	if bookingErr != nil {
		log.Println(bookingErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Finish is API for finishing booking that marshaled from request body
func Finish(w http.ResponseWriter, r *http.Request) {

	u, err := readRequestBody[entity.Booking](r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, _ := api.RetrieveRefreshToken(r)
	bookingErr := kernel.FinishBooking(u, rest.New(token))

	if bookingErr != nil {
		log.Println(bookingErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func readRequestBody[T any](r *http.Request) (T, error) {
	var res T
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return res, err
	}
	log.Printf("reading request %s", body)
	unmarshalErr := json.Unmarshal(body, &res)
	if unmarshalErr != nil {
		log.Println(unmarshalErr)
		return res, unmarshalErr
	}

	log.Printf("reading request unmarshaled into %v", res)
	return res, nil
}
