package api

import (
	"booking-service/internal/entity"
	"booking-service/internal/kernel/watcher"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u entity.Booking
	unmarshalErr := json.Unmarshal(body, &u)
	if unmarshalErr != nil {
		log.Println(unmarshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bookingID, createErr := watcher.Create(u, watcher.GetInstance())
	if createErr != nil {
		log.Println(createErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u.ID = bookingID

	responseBody, marshalErr := json.Marshal(u)
	if marshalErr != nil {
		log.Println(marshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)

}

func ChooseBooking(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u entity.Booking
	unmarshalErr := json.Unmarshal(body, &u)
	if unmarshalErr != nil {
		log.Println(unmarshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if u.CarID == 0 {
		log.Println("car must be set")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("writing booking %v", u)
	chooseErr := watcher.Choose(u, watcher.GetInstance())
	if chooseErr != nil {
		log.Println(chooseErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CancelBooking(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u entity.Booking
	unmarshalErr := json.Unmarshal(body, &u)
	if unmarshalErr != nil {
		log.Println(unmarshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chooseErr := watcher.Cancel(u, watcher.GetInstance())
	if chooseErr != nil {
		log.Println(chooseErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ApproveBooking(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u entity.Booking
	unmarshalErr := json.Unmarshal(body, &u)
	if unmarshalErr != nil {
		log.Println(unmarshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chooseErr := watcher.Approve(u, watcher.GetInstance())
	if chooseErr != nil {
		log.Println(chooseErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
