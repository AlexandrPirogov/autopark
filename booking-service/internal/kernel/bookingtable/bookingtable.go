// btable is hash-table data structure that collects userID and respective BookingFSM
package btable

import (
	"booking-service/internal/entity"
	"booking-service/internal/kernel/bookingfsm"
	"fmt"
	"log"
	"sync"
)

// New creates and returns pointer to the new instance of btable
func New() *btable {
	return &btable{
		mutex: sync.RWMutex{},
		fsms:  make(map[int]*bookingfsm.BookingFSM),
	}
}

// btable struct that wraps map and mutes agains race condition
type btable struct {
	mutex sync.RWMutex
	fsms  map[int]*bookingfsm.BookingFSM
}

// Create creates new record in btable
//
// Pre-cond: given entity with set userID
//
// Post-cond: if instance doesn't exists in btable -- insert it to the table and returns nil.
// Otherwise returns error.
func (bt *btable) Create(b entity.Booking) error {
	if _, ok := bt.fsms[b.UserID]; ok {
		return fmt.Errorf("user already created booking")
	}

	bt.fsms[b.UserID] = bookingfsm.New(b.ID, b.UserID)
	return nil
}

// Choose transform bookingFSM to StateChoosedCar
//
// Pre-cond: given entity.Boking with set userID and CarID
//
// Post-cond: if respective record exists -- trying to make
// transformation and return result of transformation.
// If record doesn't exists then returns error
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

// Approve transform bookingFSM to StateApproved
//
// Pre-cond: given entity.Boking with set userID
//
// Post-cond: if respective record exists -- trying to make
// transformation and return result of transformation.
// If record doesn't exists then returns error
func (bt *btable) Approve(b entity.Booking) error {
	bt.mutex.RLock()
	defer bt.mutex.RUnlock()
	log.Printf("booking %v was approved", b)
	if val, ok := bt.fsms[b.UserID]; ok {
		bt.fsms[b.UserID].Book = b

		err := val.Approve()
		if err != nil {
			log.Println("err while approving booking: ", err)
			return err
		}

		log.Printf("deleting %v", b)
		delete(bt.fsms, b.UserID)
	}
	return nil
}

// Cancel transform bookingFSM to StateCanceled
//
// Pre-cond: given entity.Boking with set userID
//
// Post-cond: if respective record exists -- trying to make
// transformation and return result of transformation.
// If transformation was successfull then deletes record from btable
// If record doesn't exists then returns error
func (bt *btable) Cancel(b entity.Booking) error {
	bt.mutex.RLock()
	defer bt.mutex.RUnlock()
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

// ExistsInTable checks if record with given entity.Booking is exists
//
// Pre-cond: given entity.Boking
//
// Post-cond: returns existence of given entity.Booking in btable
func (bt *btable) ExistsInTable(b entity.Booking) bool {
	bt.mutex.RLock()
	defer bt.mutex.RUnlock()
	log.Printf("checking booking %v ", b)
	_, ok := bt.fsms[b.UserID]
	return ok
}
