package postgres

// QueryLookForAdmin reads id of admin by given login and pwd
const QueryLookForAdmin = `select id from admins where login = $1 and pwd = $2`

// QueryLookForManager reads id of manager by given login and pwd
const QueryLookForManager = `select id from managers where login = $1 and pwd = $2`

// QueryInsertNewManager inserts new manager with given login and pwd.
const QueryInsertNewManager = `insert into managers values(default, $1, $2)`

// QueryLookForClients reads id of manager by given login and pwd
const QueryLookForClient = `select id from clients where login = $1 and pwd = $2`

// QueryInsertNewClient inserts new manager with given login and pwd.
const QueryInsertNewClient = `insert into clients values(default, $1, $2)`
