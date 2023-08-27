package postgres

const QueryLookForAdmin = `select id from admins where login = $1 and pwd = $2`

const QueryLookForManager = `select id from managers where login = $1 and pwd = $2`

const QueryInsertNewManager = `insert into managers values(default, $1, $2)`
