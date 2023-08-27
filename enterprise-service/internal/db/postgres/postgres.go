// package postgres hold functionality to work with Postgres
package postgres

import (
	"context"
	"enterprise-service/internal/client"
	"enterprise-service/internal/enterprise"
	"enterprise-service/internal/std"
	"enterprise-service/internal/std/list"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Instance of singleton
var singletonConn *pgconn = nil
var initOnce sync.Once

// GetInstance returns singleton instance of postgres connection
//
// Pre-cond:
//
// Post-cond: returns pointer to singleton postgres connection
func GetInstance() *pgconn {
	initOnce.Do(func() {
		singletonConn = new()
	})
	return singletonConn
}

type pgconn struct {
	conn *pgxpool.Pool
}

// AssignManager assignes given manager to enterprise
//
// Pre-cond: given Manager to assign with set ID and E_ID
//
// Post-cond: EnterpriseStorer assign given manager to enterprise.
// If successfull, returns nil otherwise returns error
func (pg *pgconn) AssignManager(m client.Manager) error {
	_, err := pg.conn.Exec(context.Background(), QueryAssignManager, m.EnterpriseID, m.Id)
	if err != nil {
		return err
	}
	return nil
}

// Delete removes all enterprise entities which mathces given pattern
//
// Pre-cond: given pattern
//
// Post-cond: all enterprises which matches given pattern was removed.
// If successfull error equals nil
func (pg *pgconn) Delete(e enterprise.Enterprise) error {
	return nil
}

// Read returns all enterprise entities from given postgres
//
// Pre-cond: given pattern
//
// Post-cond: all enterprises which matches given pattern was returned.
// If successfull error equals nil, otherwise returns error
func (pg *pgconn) Read(e enterprise.Enterprise) (std.Linked[enterprise.Enterprise], error) {
	rows, err := pg.conn.Query(context.Background(), QueryReadEnterprises)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := list.New[enterprise.Enterprise]()
	for rows.Next() {
		var e enterprise.Enterprise
		scanErr := rows.Scan(&e.Title)
		if scanErr != nil {
			log.Println(err)
			continue
		}
		res.PushBack(e)
	}

	return res, nil
}

// ReadByID returns enterprise entity with given id
//
// Pre-cond: given positive id
//
// Post-cond: returnes enterprise with given ID from EnterpriseStorer
func (pg *pgconn) ReadByID(id int) (enterprise.Enterprise, error) {
	var e enterprise.Enterprise
	err := pg.conn.QueryRow(context.Background(), QueryReadByIDEnterprises, id).Scan(&e.Title)
	return e, err
}

// StoreEnterprise writes given enterprise entity
//
// Pre-cond: given new unique enterprise entity to write
//
// Post-cond: given enterprise entity was written
func (pg *pgconn) StoreEnterprise(e enterprise.Enterprise) error {
	_, err := pg.conn.Exec(context.Background(), QueryStoreEnterprise, e.Title)
	if err != nil {
		return err
	}
	return nil
}

// new returnes new instance of pgconn
func new() *pgconn {
	return &pgconn{
		conn: tryConnect(),
	}
}

// tryConnect tryies to connect to postgres
// if failed throws an exeption
func tryConnect() *pgxpool.Pool {
	URL := "postgresql://postgres:postgres@enterprise-db:5432/postgres"
	conn, err := pgxpool.New(context.Background(), URL)
	if err != nil {
		log.Fatal(err)
	}

	ping := conn.Ping(context.Background())
	if ping != nil {
		log.Fatalf("error while pinging %v", ping)
	}
	return conn
}
