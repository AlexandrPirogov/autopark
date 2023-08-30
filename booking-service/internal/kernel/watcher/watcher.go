package watcher

import (
	"booking-service/internal/db"
	"booking-service/internal/entity"
	btable "booking-service/internal/kernel/bookingtable"
	"log"
	"sync"
	"time"
)

type BookingTabler interface {
	Approve(b entity.Booking) error
	Cancel(b entity.Booking) error
	Choose(b entity.Booking) error
	Create(b entity.Booking) error
}

type watcher struct {
	s     db.BookingStorer
	table BookingTabler
}

var singletonWatcher *watcher

var initOnce sync.Once

// GetInstance returns *pgconn
// If pgconn was btable returns the old one
func GetInstance() *watcher {
	initOnce.Do(func() {
		singletonWatcher = new()
	})
	return singletonWatcher
}

func new() *watcher {
	return &watcher{
		s:     db.GetConnInstance(),
		table: btable.New(),
	}
}

func Approve(b entity.Booking, w *watcher) error {
	tableErr := w.table.Approve(b)
	if tableErr != nil {
		log.Println(tableErr)
		return tableErr
	}

	dbErr := w.s.Approved(b)
	if dbErr != nil {
		log.Println(dbErr)
		return dbErr
	}

	return nil
}

func Cancel(b entity.Booking, w *watcher) error {
	tableErr := w.table.Cancel(b)
	if tableErr != nil {
		log.Println(tableErr)
		return tableErr
	}

	dbErr := w.s.Canceled(b)
	if dbErr != nil {
		log.Println(dbErr)
		return dbErr
	}

	return nil
}

func Choose(b entity.Booking, w *watcher) error {
	tableErr := w.table.Choose(b)
	if tableErr != nil {
		log.Println(tableErr)
		return tableErr
	}

	dbErr := w.s.CarChoosed(b)
	if dbErr != nil {
		log.Println(dbErr)
		return dbErr
	}

	return nil
}

func Create(b entity.Booking, w *watcher) (int, error) {
	go func() {
		time.AfterFunc(time.Minute*1, func() {
			tableErr := w.table.Cancel(b)
			if tableErr == nil {
				w.s.Canceled(b)
			}
		})
	}()

	id, dbErr := w.s.Created(b)
	if dbErr != nil {
		log.Println(dbErr)
		return -1, dbErr
	}

	b.ID = id
	tableErr := w.table.Create(b)
	if tableErr != nil {
		log.Println(dbErr)
		return -1, dbErr
	}

	return id, nil
}
