// bookingfsm represents FiniteStateMachine for booking.
// Book can be on on the next states: created, choosed, approved ,cancel, finished.
// Each of these states described below.
package bookingfsm

import (
	"booking-service/internal/entity"
	"context"
	"fmt"
	"log"

	"github.com/looplab/fsm"
)

const (
	StateCreated    = "created"  // Booking was created
	StateChoosedCar = "choosed"  // Car for created booking was choosed. Can be performed only from StateCreated
	StateApproved   = "approved" // Booking was approved. Can pe performed only from StateChoosedCar
	StateCanceled   = "canceled" // Booking was canceled. Can be performed from StateCreated and StateChoosedCar
	StateFinished   = "finished" // Booking was finished. Can be performed only from StateApproved
)

// BookingFSM is a wrapper for looplabfsm
type BookingFSM struct {
	Book entity.Booking
	fsm  *fsm.FSM
}

// Created new BooingFSM instance
//
// Pre-cond: given id and userID
//
// Post-cond: instance was created and pointer to it was returned
func New(id, userID int) *BookingFSM {
	return &BookingFSM{
		Book: entity.Booking{
			ID:     id,
			UserID: userID,
		},
		fsm: fsm.NewFSM(
			StateCreated,
			fsm.Events{
				{Name: StateCreated, Src: []string{StateCreated}, Dst: StateCreated},
				{Name: StateChoosedCar, Src: []string{StateCreated}, Dst: StateChoosedCar},
				{Name: StateApproved, Src: []string{StateChoosedCar}, Dst: StateApproved},
				{Name: StateCanceled, Src: []string{StateCreated, StateChoosedCar}, Dst: StateCanceled},
				{Name: StateFinished, Src: []string{StateApproved}, Dst: StateFinished},
			},
			fsm.Callbacks{},
		),
	}
}

// Current return current state of BookingFSM
func (b *BookingFSM) Current() string {
	return b.fsm.Current()
}

// ChooseCar tries to perfom transformation to StateChoosedCar.
// If transformation was performed successfully returns nil, otherwise returns error
func (b *BookingFSM) ChooseCar() error {
	err := b.fsm.Event(context.Background(), StateChoosedCar)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Printf("Booking with id %d has choosen car. Current fsm state: %s\n", b.Book.ID, b.fsm.Current())
	return nil
}

// Approve tries to perfom transformation to StateApproved
// If transformation was performed successfully returns nil, otherwise returns error
func (b *BookingFSM) Approve() error {
	err := b.fsm.Event(context.Background(), StateApproved)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Printf("Booking with id %d has approved. Current fsm state: %s\n", b.Book.ID, b.fsm.Current())
	return nil
}

// Cancel tries to perfom transformation to StateCanceled.
// If transformation was performed successfully returns nil, otherwise returns error
func (b *BookingFSM) Cancel() error {
	err := b.fsm.Event(context.Background(), StateCanceled)
	if err != nil {
		log.Println("cancel event called ", err)
		return err
	}
	fmt.Printf("Booking with id %d has canceled. Current fsm state: %s\n", b.Book.ID, b.fsm.Current())
	return nil
}

// Finish tries to perfom transformation to StateFinished.
// If transformation was performed successfully returns nil, otherwise returns error
func (b *BookingFSM) Finish() error {
	err := b.fsm.Event(context.Background(), StateFinished)
	if err != nil {
		return err
	}
	fmt.Printf("Booking with id %d has finished. Current fsm state: %s\n", b.Book.ID, b.fsm.Current())
	return nil
}

// Booking returns copy of entity.Booking
func (b *BookingFSM) Booking() entity.Booking {
	return b.Book
}
