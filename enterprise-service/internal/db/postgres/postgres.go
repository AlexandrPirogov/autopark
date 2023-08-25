package postgres

import (
	"enterprise-service/internal/enterprise"
	"enterprise-service/internal/std"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Instance of singleton
var singletonConn *pgconn = nil
var initOnce sync.Once

func GetInstance() *pgconn {
	initOnce.Do(func() {
		singletonConn = new()
	})
	return singletonConn
}

type pgconn struct {
	conn *pgxpool.Pool
}

func new() *pgconn {
	return &pgconn{}
}

func (pg *pgconn) Delete(e enterprise.Enterprise) error {
	return nil
}

func (pg *pgconn) Read(e enterprise.Enterprise) (std.Linked[enterprise.Enterprise], error) {
	return nil, nil
}

func (pg *pgconn) Store(e enterprise.Enterprise) error {
	return nil
}

func (pg *pgconn) Update(s enterprise.Enterprise) (std.Linked[enterprise.Enterprise], error) {
	return nil, nil
}
