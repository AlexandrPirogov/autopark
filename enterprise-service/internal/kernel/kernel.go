// package kernel collects main functionality for enterprise
package kernel

import (
	"enterprise-service/internal/enterprise"
	"enterprise-service/internal/std"
)

// EnterpriseStorer stands for working with storer to save enterprise entyties
type EntepriseStorer interface {
	Delete(e enterprise.Enterprise, s EntepriseStorer) error
	Read(e enterprise.Enterprise, s EntepriseStorer) (std.Linked[enterprise.Enterprise], error)
	Store(e enterprise.Enterprise, s EntepriseStorer) error
	Update(e EntepriseStorer, s EntepriseStorer) (std.Linked[enterprise.Enterprise], error)
}
