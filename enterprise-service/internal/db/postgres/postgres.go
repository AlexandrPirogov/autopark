// package postgres hold functionality to work with Postgres
package postgres

import (
	"context"
	"enterprise-service/internal/client"
	"enterprise-service/internal/config"
	"enterprise-service/internal/enterprise"
	"enterprise-service/internal/std"
	"enterprise-service/internal/std/list"
	"sync"

	"github.com/rs/zerolog/log"

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
	log.Warn().Msgf("assigning manager %v to postgres", m)
	_, err := pg.conn.Exec(context.Background(), QueryAssignManager, m.EnterpriseTitle, m.Name, m.Surname)
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
func (pg *pgconn) Read() (std.Linked[enterprise.Enterprise], error) {
	rows, err := pg.conn.Query(context.Background(), QueryReadEnterprises)
	if err != nil {
		log.Warn().Msgf("%v", err)
		return nil, err
	}
	defer rows.Close()
	res := list.New[enterprise.Enterprise]()
	for rows.Next() {
		var e enterprise.Enterprise
		scanErr := rows.Scan(&e.ID, &e.Title)
		if scanErr != nil {
			log.Warn().Msgf("%v", err)
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
func (pg *pgconn) ReadByTitle(title string) (enterprise.Enterprise, error) {
	var e enterprise.Enterprise
	err := pg.conn.QueryRow(context.Background(), QueryReadByTitleEnterprises, title).Scan(&e.ID, &e.Title)
	if err != nil {
		log.Warn().Msgf("err while executing read by title %v", err)
		return e, err
	}

	rows, err := pg.conn.Query(context.Background(), QueryReadByTitleEnterprisesManagers, title)
	if err != nil {
		log.Warn().Msgf("err while executing read by title %v", err)
		return e, err
	}

	defer rows.Close()
	for rows.Next() {
		var m client.Manager
		scanErr := rows.Scan(&m.Name, &m.Surname)
		if scanErr != nil {
			log.Warn().Msgf("err while scanning managers%v", err)
			continue
		}

		e.Managers = append(e.Managers, m)
	}
	log.Warn().Msgf("read by title result %v", e)
	return e, err
}

// StoreEnterprise writes given enterprise entity
//
// Pre-cond: given new unique enterprise entity to write
//
// Post-cond: given enterprise entity was written
func (pg *pgconn) StoreEnterprise(e enterprise.Enterprise) error {
	log.Warn().Msgf("registering %v", e)
	_, err := pg.conn.Exec(context.Background(), QueryStoreEnterprise, e.Title)
	log.Warn().Msgf("registering err %v", err)
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
	URL := config.PostgresURL()
	conn, err := pgxpool.New(context.Background(), URL)
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}

	ping := conn.Ping(context.Background())
	if ping != nil {
		log.Fatal().Msgf("error while pinging %v", ping)
	}
	return conn
}
