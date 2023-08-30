package bookingfsm

import (
	"booking-service/internal/entity"
	"context"
	"fmt"
	"log"

	"github.com/looplab/fsm"
)

const (
	StateCreated    = "created"
	StateChoosedCar = "choosed"
	StateApproved   = "approved"
	StateCanceled   = "canceled"
)

type BookingFSM struct {
	Book entity.Booking
	fsm  *fsm.FSM
}

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
			},
			fsm.Callbacks{},
		),
	}
}

func (b *BookingFSM) Current() string {
	return b.fsm.Current()
}

func (b *BookingFSM) ChooseCar() error {
	err := b.fsm.Event(context.Background(), StateChoosedCar)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Printf("Booking with id %d has choosen car. Current fsm state: %s\n", b.Book.ID, b.fsm.Current())
	return nil
}

func (b *BookingFSM) Approve() error {
	err := b.fsm.Event(context.Background(), StateApproved)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Printf("Booking with id %d has approved. Current fsm state: %s\n", b.Book.ID, b.fsm.Current())
	return nil
}

func (b *BookingFSM) Cancel() error {
	err := b.fsm.Event(context.Background(), StateCanceled)
	if err != nil {
		log.Println("cancel event called ", err)
		return err
	}
	fmt.Printf("Booking with id %d has canceled. Current fsm state: %s\n", b.Book.ID, b.fsm.Current())
	return nil
}

func (b *BookingFSM) Booking() entity.Booking {
	return b.Book
}
