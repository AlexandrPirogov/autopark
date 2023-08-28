package postgres

// QueryStoreEnterprise stores given Enterprise. Enterprise's title must be unique
const QueryStoreEnterprise = "insert into enterprises values(default, $1)"

// QueryReadEnterprises reads all enterprises from storage
const QueryReadEnterprises = "select title from enterprises"

// QueryReadByIDEnterprises return enterprise with given ID
const QueryReadByIDEnterprises = "select title from enterprises where id = $1"

// QueryAssignManager assign given manager to given enterprise
const QueryAssignManager = "insert into ENTERPRISES_MANAGERS values($1, $2)"
