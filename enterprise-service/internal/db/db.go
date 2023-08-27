package db

import (
	"enterprise-service/internal/client"
	"enterprise-service/internal/db/postgres"
	"enterprise-service/internal/enterprise"
	"enterprise-service/internal/std"
)

// EnterpriseStorer stands for working with storer to save enterprise entyties
type EntepriseStorer interface {
	Delete(e enterprise.Enterprise) error
	Read(e enterprise.Enterprise) (std.Linked[enterprise.Enterprise], error)
	ReadByID(id int) (enterprise.Enterprise, error)
	StoreEnterprise(e enterprise.Enterprise) error
	AssignManager(m client.Manager) error
	Update(e enterprise.Enterprise) (std.Linked[enterprise.Enterprise], error)
}

func GetConnInstance() EntepriseStorer {
	return postgres.GetInstance()
}
