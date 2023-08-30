// watcher wrapper aroung in memory btable and db connection
package watcher

import (
	"booking-service/internal/db"
	"booking-service/internal/entity"
	btable "booking-service/internal/kernel/bookingtable"
	"fmt"
	"log"
	"sync"
	"time"
)

// Interface for btable
type BookingTabler interface {
	Approve(entity.Booking) error
	Cancel(entity.Booking) error
	Choose(entity.Booking) error
	Create(entity.Booking) error
	ExistsInTable(entity.Booking) bool
}

type watcher struct {
	s     db.BookingStorer
	table BookingTabler
}

// singletonWatcher is signleton
var singletonWatcher *watcher
var initOnce sync.Once

// GetInstance returns *watcher
// If watcher was created returns pointer to it
func GetInstance() *watcher {
	initOnce.Do(func() {
		singletonWatcher = new()
	})
	return singletonWatcher
}

// new returns pointer to the singleton of wather
func new() *watcher {
	return &watcher{
		s:     db.GetConnInstance(),
		table: btable.New(),
	}
}

// Approve updates btable's record for given entity.Booking
// and update respective record in database/
//
// Pre-cond: given entity.Booking to approve and watcher.
// Updating entity.Booking must exists in table
//
// Post-cond: if transformation was executed successfully
// the writes updates to database. First it's searching record in btable.
// If transformation was failed then cancel the booking
func Approve(b entity.Booking, w *watcher) error {
	if !w.table.ExistsInTable(b) {
		log.Println("cant approve not existing booking")
		return fmt.Errorf("cant approve not existing booking")
	}

	tableErr := w.table.Approve(b)
	if tableErr != nil {
		log.Println("btable err ", tableErr)
		return tableErr
	}

	dbErr := w.s.Approved(b)
	if dbErr != nil {
		log.Println("db err ", dbErr)
		return dbErr
	}

	return nil
}

// Cancel updates btable's record for given entity.Booking
// and update respective record in database/
//
// Pre-cond: given entity.Booking to Cancel and watcher.
// Updating entity.Booking must exists in table
//
// Post-cond: if transformation was executed successfully
// the writes updates to database. First it's searching record in btable.
// If transformation was failed then cancel the booking
func Cancel(b entity.Booking, w *watcher) error {

	if !w.table.ExistsInTable(b) {
		log.Println("cant cancel not existing booking")
		return fmt.Errorf("cant cancel not existing booking")
	}

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

// Choose updates btable's record for given entity.Booking
// and update respective record in database/
//
// Pre-cond: given entity.Booking to Choose and watcher.
// Updating entity.Booking must exists in table
//
// Post-cond: if transformation was executed successfully
// the writes updates to database. First it's searching record in btable.
// If transformation was failed then cancel the booking
func Choose(b entity.Booking, w *watcher) error {

	if !w.table.ExistsInTable(b) {
		log.Println("cant Choose not existing booking")
		return fmt.Errorf("cant Choose not existing booking")
	}

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

// Create updates btable's record for given entity.Booking
// and update respective record in database/
//
// Pre-cond: given entity.Booking to Create and watcher.
// Updating entity.Booking must not exists in table
//
// Post-cond: if entity.Booking exists in table then returns error
// Otherwise given entity.Booking was added to btable and record
// was added to database
func Create(b entity.Booking, w *watcher) (int, error) {

	if w.table.ExistsInTable(b) {
		return -1, fmt.Errorf("booking already created")
	}

	id, dbErr := w.s.Created(b)
	if dbErr != nil {
		log.Println(dbErr)
		return -1, dbErr
	}

	b.ID = id
	tableErr := w.table.Create(b)
	if tableErr != nil {
		log.Println(tableErr)
		return -1, tableErr
	}
	go func() {
		time.AfterFunc(time.Minute*1, func() {
			tableErr := w.table.Cancel(b)
			if tableErr == nil {
				w.s.Canceled(b)
			}
		})
	}()
	return id, nil
}

// Finish updates btable's record for given entity.Booking
// and update respective record in database/
//
// Pre-cond: given entity.Booking to Finish and watcher.
// Updating entity.Booking must not exists in table
//
// Post-cond: if transformation was executed successfully
// the writes updates to database. The records MUST NOT EXISTS in btable.
// If it exists returns error otherwise returnes the result of transforamtion.
func Finish(b entity.Booking, w *watcher) error {
	if w.table.ExistsInTable(b) {
		log.Println("cant finish existing booking")
		return fmt.Errorf("cant finish existing booking")
	}

	dbErr := w.s.Finish(b)
	if dbErr != nil {
		log.Println(dbErr)
		return dbErr
	}

	return nil
}
