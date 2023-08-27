package postgres

const QueryLookForAdmin = `select uid from admins where login = $1 and pwd = $2`
