// package kernel collects main functionality for enterprise
package kernel

import (
	"enterprise-service/internal/enterprise"
	"enterprise-service/internal/std"
)

// EnterpriseStorer stands for working with storer to save enterprise entyties
type EntepriseStorer interface {
	Delete(e enterprise.Enterprise) error
	Read(e enterprise.Enterprise) (std.Linked[enterprise.Enterprise], error)
	Store(e enterprise.Enterprise) error
	Update(e enterprise.Enterprise) (std.Linked[enterprise.Enterprise], error)
}

// Delete removes all enterprise entities from given EnterpriseStore which mathces given pattern
//
// Pre-cond: given pattern and Storer to remove enterprise entityies from
//
// Post-cond: all enterprises which matches given pattern was removed from given EnterpriseStorer
func Delete(pattern enterprise.Enterprise, s EntepriseStorer) error {
	return s.Delete(pattern)
}

// Read returns all enterprise entities from given EnterpriseStore which mathces given pattern
//
// Pre-cond: given pattern and Storer to returns enterprise entities from
//
// Post-cond: all enterprises which matches given pattern was returned from given EnterpriseStorer
func Read(pattern enterprise.Enterprise, s EntepriseStorer) (std.Linked[enterprise.Enterprise], error) {
	return s.Read(pattern)
}

// Store writes given enterprise entity to given EnterpriseStore
//
// Pre-cond: given new unique enterprise entity to write and EnterpriseStore to write to
//
// Post-cond: given enterprise entity was written to EnterpriseStore
func Store(e enterprise.Enterprise, s EntepriseStorer) error {
	return s.Store(e)
}

// Update updates all enterprise entities from given EnterpriseStore which mathces given pattern
//
// Pre-cond: given pattern and EnterpriseStorer to update enterprise entities
//
// Post-cond: all enterprises which matches given pattern was updated within given EnterpriseStorer
func Update(pattern enterprise.Enterprise, s EntepriseStorer) (std.Linked[enterprise.Enterprise], error) {
	return s.Update(pattern)
}
