package btable

import (
	"booking-service/internal/entity"
	"booking-service/internal/kernel/bookingfsm"
	"fmt"
	"log"
	"sync"
)

// new creates new connection to Postgres via pool
func New() *btable {
	return &btable{
		mutex: sync.RWMutex{},
		fsms:  make(map[int]*bookingfsm.BookingFSM),
	}
}

type btable struct {
	mutex sync.RWMutex
	fsms  map[int]*bookingfsm.BookingFSM
}

func (bt *btable) Create(b entity.Booking) error {
	if _, ok := bt.fsms[b.UserID]; ok {
		return fmt.Errorf("user already created booking")
	}

	bt.fsms[b.UserID] = bookingfsm.New(b.ID, b.UserID)
	return nil
}

func (bt *btable) Choose(b entity.Booking) error {
	bt.mutex.RLock()
	defer bt.mutex.RUnlock()
	if val, ok := bt.fsms[b.UserID]; ok {
		bt.fsms[b.UserID].Book = b

		err := val.ChooseCar()
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (bt *btable) Approve(b entity.Booking) error {
	bt.mutex.RLock()
	defer bt.mutex.RUnlock()
	log.Printf("booking %v was approved", b)
	if val, ok := bt.fsms[b.UserID]; ok {
		bt.fsms[b.UserID].Book = b

		err := val.Approve()
		if err != nil {
			log.Println(err)
			return err
		}

		log.Printf("deleting %v", b)
		delete(bt.fsms, b.UserID)
	}
	return nil
}

func (bt *btable) Cancel(b entity.Booking) error {
	bt.mutex.RLock()
	defer bt.mutex.RUnlock()
	log.Printf("timeout for %v", b)
	if val, ok := bt.fsms[b.UserID]; ok {
		err := val.Cancel()
		if err != nil {
			log.Println(err)
			return err
		}
		log.Printf("deleting %v", b)
		delete(bt.fsms, b.UserID)
		return nil
	}
	return fmt.Errorf("not found")
}
