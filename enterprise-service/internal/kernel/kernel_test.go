package kernel

import (
	"enterprise-service/internal/enterprise"
	"enterprise-service/internal/std"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dbstub struct{}

func (db *dbstub) Delete(e enterprise.Enterprise) error {
	return nil
}

func (db *dbstub) Read(e enterprise.Enterprise) (std.Linked[enterprise.Enterprise], error) {
	return nil, nil
}

func (db *dbstub) Store(e enterprise.Enterprise) error {
	return nil
}

func (db *dbstub) Update(s enterprise.Enterprise) (std.Linked[enterprise.Enterprise], error) {
	return nil, nil
}

func TestStore(t *testing.T) {
	stub := &dbstub{}
	e := enterprise.Enterprise{}

	err := Store(e, stub)
	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	stub := &dbstub{}
	e := enterprise.Enterprise{}

	_, err := Update(e, stub)
	assert.Nil(t, err)
}

func TestRead(t *testing.T) {
	stub := &dbstub{}
	e := enterprise.Enterprise{}

	_, err := Read(e, stub)
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	stub := &dbstub{}
	e := enterprise.Enterprise{}

	err := Delete(e, stub)
	assert.Nil(t, err)
}
