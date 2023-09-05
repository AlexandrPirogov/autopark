package postgres

// QueryStoreEnterprise stores given Enterprise. Enterprise's title must be unique
const QueryStoreEnterprise = "insert into enterprises values(default, $1)"

// QueryReadEnterprises reads all enterprises from storage
const QueryReadEnterprises = "select id, title from enterprises"

// QueryReadByIDEnterprises return enterprise with given ID
const QueryReadByIDEnterprises = "select id, title from enterprises where id = $1"

// QueryReadByIDEnterprises return enterprise with given ID
const QueryReadByTitleEnterprises = `select e.id, e.title from enterprises e where e.title = $1`

// QueryReadByIDEnterprises return enterprise with given ID
const QueryReadByTitleEnterprisesManagers = `select m.name, m.surname
from enterprises e
join managers m on m.e_id = e.id and e.title = $1`

// QueryAssignManager assign given manager to given enterprise
const QueryAssignManager = `insert into managers values(default, 
		(select id from enterprises where title = $1), $2, $3)`
