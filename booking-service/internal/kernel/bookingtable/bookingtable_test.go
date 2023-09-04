package btable

import (
	"booking-service/internal/entity"
	"booking-service/internal/kernel/bookingfsm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBooking(t *testing.T) {
	sut := New()

	sut.Create(entity.Booking{UserID: 25})
	assert.Equal(t, len(sut.fsms), 1)
	assert.Equal(t, bookingfsm.StateCreated, sut.fsms[25].Current())
}

func TestCreateTwiceBooking(t *testing.T) {
	sut := New()

	sut.Create(entity.Booking{UserID: 25})
	err := sut.Create(entity.Booking{UserID: 25})
	assert.NotNil(t, err)
	assert.Equal(t, len(sut.fsms), 1)
	assert.Equal(t, bookingfsm.StateCreated, sut.fsms[25].Current())
}

func TestCancelExisting(t *testing.T) {
	b := entity.Booking{UserID: 25}
	sut := New()
	sut.Create(b)
	old := len(sut.fsms)
	sut.Cancel(b)
	new := len(sut.fsms)

	assert.Greater(t, old, new)
}

func TestCreateAndChoose(t *testing.T) {
	b := entity.Booking{UserID: 25}
	sut := New()
	sut.Create(b)
	sut.Choose(b)

	assert.Equal(t, bookingfsm.StateChoosedCar, sut.fsms[b.UserID].Current())
}

func TestCreateAndChooseAndApprove(t *testing.T) {
	b := entity.Booking{UserID: 25}
	sut := New()
	sut.Create(b)
	sut.Choose(b)
	sut.Approve(b)

	assert.Equal(t, bookingfsm.StateApproved, sut.fsms[b.UserID].Current())
}
