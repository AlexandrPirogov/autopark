// package kernel collects main functionality for enterprise
package kernel

import (
	"enterprise-service/internal/client"
	"enterprise-service/internal/db"
	"enterprise-service/internal/enterprise"
	"enterprise-service/internal/std"
)

// RegisterManager sends request by ManagerHandler to register new manager in system
//
// Pre-cond: given Manage instance with login and pwd and ManagerHandler to send request
//
// Post-cond: request was executed. If manager was registered successfully,
// returns manager with set ID and nil. Otherwise return error
func RegisterManager(m client.Manager, c client.ManagerHandler) (client.Manager, error) {
	return c.RegisterManager(m)
}

// AssignManager assignes registered manager to enterprise in enterprise-system
//
// Pre-cond: given Manager to assign with set ID and E_ID and EnterpriseStorer to store in
//
// Post-cond: EnterpriseStorer assign given manager to enterprise. If successfull, returns nil
// otherwise returns error
func AssignManager(m client.Manager, s db.EntepriseStorer) error {
	return s.AssignManager(m)
}

// Delete removes all enterprise entities from given EnterpriseStore which mathces given pattern
//
// Pre-cond: given pattern and Storer to remove enterprise entityies from
//
// Post-cond: all enterprises which matches given pattern was removed from given EntepriseStorer
func Delete(pattern enterprise.Enterprise, s db.EntepriseStorer) error {
	return s.Delete(pattern)
}

// Read returns all enterprise entities from given EnterpriseStore which mathces given pattern
//
// Pre-cond: given pattern and Storer to returns enterprise entities from
//
// Post-cond: all enterprises which matches given pattern was returned from given EntepriseStorer
func Read(pattern enterprise.Enterprise, s db.EntepriseStorer) (std.Linked[enterprise.Enterprise], error) {
	return s.Read(pattern)
}

// ReadByID returns enterprise entity with given id from given EnterpriseStore which mathces given pattern
//
// Pre-cond: given positive id and Storer to returns terprise entities from
//
// Post-cond: returnes enterprise with given ID from EnterpriseStorer
func ReadByID(id int, s db.EntepriseStorer) (enterprise.Enterprise, error) {
	return s.ReadByID(id)
}

// Store writes given enterprise entity to given EnterpriseStore
//
// Pre-cond: given new unique enterprise entity to write and EnterpriseStore to write to
//
// Post-cond: given enterprise entity was written to EnterpriseStore
func Store(e enterprise.Enterprise, s db.EntepriseStorer) error {
	return s.StoreEnterprise(e)
}

// Update updates all enterprise entities from given EnterpriseStore which mathces given pattern
//
// Pre-cond: given pattern and EntepriseStorer to update enterprise entities
//
// Post-cond: all enterprises which matches given pattern was updated within given EntepriseStorer
func Update(pattern enterprise.Enterprise, s db.EntepriseStorer) (std.Linked[enterprise.Enterprise], error) {
	return s.Update(pattern)
}
