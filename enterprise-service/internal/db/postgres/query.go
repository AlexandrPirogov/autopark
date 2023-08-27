package postgres

const QueryStoreEnterprise = "insert into enterprises values(default, $1)"

const QueryReadEnterprises = "select title from enterprises"

const QueryReadByIDEnterprises = "select title from enterprises where id = $1"
