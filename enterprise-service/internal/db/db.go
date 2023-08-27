package db

import (
	"enterprise-service/internal/db/postgres"
	"enterprise-service/internal/enterprise"
	"enterprise-service/internal/std"
)

// EnterpriseStorer stands for working with storer to save enterprise entyties
type EntepriseStorer interface {
	Delete(e enterprise.Enterprise) error
	Read(e enterprise.Enterprise) (std.Linked[enterprise.Enterprise], error)
	ReadByID(id int) (enterprise.Enterprise, error)
	Store(e enterprise.Enterprise) error
	Update(e enterprise.Enterprise) (std.Linked[enterprise.Enterprise], error)
}

func GetConnInstance() EntepriseStorer {
	return postgres.GetInstance()
}
