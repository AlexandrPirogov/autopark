// package db holds interface to store entities in enterprise-service
package db

import (
	"enterprise-service/internal/client"
	"enterprise-service/internal/db/postgres"
	"enterprise-service/internal/enterprise"
	"enterprise-service/internal/std"
)

// EnterpriseStorer stands for working with storer to save enterprise entyties
type EntepriseStorer interface {
	// AssignManager assignes given manager to enterprise
	//
	// Pre-cond: given Manager to assign with set ID and E_ID
	//
	// Post-cond: EnterpriseStorer assign given manager to enterprise.
	// If successfull, returns nil otherwise returns error
	AssignManager(m client.Manager) error

	// Delete removes all enterprise entities which mathces given pattern
	//
	// Pre-cond: given pattern
	//
	// Post-cond: all enterprises which matches given pattern was removed.
	// If successfull error equals nil
	Delete(e enterprise.Enterprise) error

	// Read returns all enterprise entities from given postgres
	//
	// Pre-cond: given pattern
	//
	// Post-cond: all enterprises which matches given pattern was returned.
	// If successfull error equals nil, otherwise returns error
	Read() (std.Linked[enterprise.Enterprise], error)

	// ReadByID returns enterprise entity with given id
	//
	// Pre-cond: given positive id
	//
	// Post-cond: returnes enterprise with given ID from EnterpriseStorer
	ReadByTitle(title string) (enterprise.Enterprise, error)

	// StoreEnterprise writes given enterprise entity
	//
	// Pre-cond: given new unique enterprise entity to write
	//
	// Post-cond: given enterprise entity was written
	StoreEnterprise(e enterprise.Enterprise) error
}

// GetConnInstance returns current connection to EnterpriseStorer
func GetConnInstance() EntepriseStorer {
	return postgres.GetInstance()
}
