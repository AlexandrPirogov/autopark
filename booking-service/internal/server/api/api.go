package api

import (
	"booking-service/internal/entity"
	"booking-service/internal/kernel/watcher"
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Warn().Msgf("%v ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u entity.Booking
	unmarshalErr := json.Unmarshal(body, &u)
	if unmarshalErr != nil {
		log.Warn().Msgf("%v ", unmarshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bookingID, createErr := watcher.Create(u, watcher.GetInstance())
	if createErr != nil {
		log.Warn().Msgf("%v", createErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u.ID = bookingID

	responseBody, marshalErr := json.Marshal(u)
	if marshalErr != nil {
		log.Warn().Msgf("%v", marshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)

}

func ChooseBooking(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Warn().Msgf("%v ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u entity.Booking
	unmarshalErr := json.Unmarshal(body, &u)
	if unmarshalErr != nil {
		log.Warn().Msgf("%v", unmarshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if u.CarID == 0 {
		log.Warn().Msgf("%v", "car must be set")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Debug().Msgf("writing booking %v", u)
	chooseErr := watcher.Choose(u, watcher.GetInstance())
	if chooseErr != nil {
		log.Warn().Msgf("%v", chooseErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CancelBooking(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Warn().Msgf("%v ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u entity.Booking
	unmarshalErr := json.Unmarshal(body, &u)
	if unmarshalErr != nil {
		log.Warn().Msgf("%v", unmarshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chooseErr := watcher.Cancel(u, watcher.GetInstance())
	if chooseErr != nil {
		log.Warn().Msgf("%v", chooseErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ApproveBooking(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Warn().Msgf("%v ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u entity.Booking
	unmarshalErr := json.Unmarshal(body, &u)
	if unmarshalErr != nil {
		log.Warn().Msgf("%v", unmarshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chooseErr := watcher.Approve(u, watcher.GetInstance())
	if chooseErr != nil {
		log.Warn().Msgf("%v", chooseErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func FinishBooking(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Warn().Msgf("%v ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u entity.Booking
	unmarshalErr := json.Unmarshal(body, &u)
	if unmarshalErr != nil {
		log.Warn().Msgf("%v", unmarshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chooseErr := watcher.Finish(u, watcher.GetInstance())
	if chooseErr != nil {
		log.Warn().Msgf("%v", chooseErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
