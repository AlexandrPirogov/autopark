package bookingfsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	sut := New(0, 0)

	assert.Equal(t, StateCreated, sut.Current())
}

func TestChooseCar(t *testing.T) {
	sut := New(0, 0)

	sut.ChooseCar()

	assert.Equal(t, StateChoosedCar, sut.Current())
}

func TestApprove(t *testing.T) {
	sut := New(0, 0)

	sut.ChooseCar()
	sut.Approve()

	assert.Equal(t, StateApproved, sut.Current())
}

func TestCancel(t *testing.T) {
	sut := New(0, 0)

	sut.Cancel()
	assert.Equal(t, StateCanceled, sut.Current())
}

func TestCancelAndCantChooseCar(t *testing.T) {
	sut := New(0, 0)

	sut.Cancel()
	assert.Equal(t, StateCanceled, sut.Current())
	assert.NotNil(t, sut.ChooseCar())
}

func TestCancelAfterChooseCar(t *testing.T) {
	sut := New(0, 0)

	sut.ChooseCar()
	sut.Cancel()

	assert.Equal(t, StateCanceled, sut.Current())
	assert.NotNil(t, sut.Approve())
}

func TestCantCancelAfterApprove(t *testing.T) {
	sut := New(0, 0)

	sut.ChooseCar()
	sut.Approve()

	assert.Equal(t, StateApproved, sut.Current())
	assert.NotNil(t, sut.Cancel())
}
