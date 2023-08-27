// package kernel collects main functionality for enterprise
package kernel

import (
	"enterprise-service/internal/db"
	"enterprise-service/internal/enterprise"
	"enterprise-service/internal/std"
)

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

// Store writes given enterprise entity to given EnterpriseStore
//
// Pre-cond: given new unique enterprise entity to write and EnterpriseStore to write to
//
// Post-cond: given enterprise entity was written to EnterpriseStore
func Store(e enterprise.Enterprise, s db.EntepriseStorer) error {
	return s.Store(e)
}

// Update updates all enterprise entities from given EnterpriseStore which mathces given pattern
//
// Pre-cond: given pattern and EntepriseStorer to update enterprise entities
//
// Post-cond: all enterprises which matches given pattern was updated within given EntepriseStorer
func Update(pattern enterprise.Enterprise, s db.EntepriseStorer) (std.Linked[enterprise.Enterprise], error) {
	return s.Update(pattern)
}
