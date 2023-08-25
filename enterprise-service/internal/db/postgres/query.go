package postgres

const QueryStoreEnterprise = "insert into enterprises values(default, $1)"

const QueryReadEnterprises = "select title from enterprises"
