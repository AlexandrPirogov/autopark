// package client hold URLs and interface ManagerHandler for requests
package client

// ApiGateway URL
const ApiGatewayHost = "http://api-gateway-nginx"

// Lists brandss
const RegisterManagerURL = "/auth/register/manager"

const RerfeshTokenCookieField = "refresh-token"

type Manager struct {
	Id              int    `json:"id"`
	EnterpriseTitle string `json:"e_title"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Login           string `json:"login"`
	Pwd             string `json:"pwd"`
}

type ManagerHandler interface {
	// RegisterManager making request to register manager in auth-service
	//
	// Pre-cond: given client.Manager instance to register
	//
	// Post-cond: request was executed and result returned.
	// If request executes successfully returns Manager that was registeres and nil error
	// Otherwise returnes nil and error
	RegisterManager(m Manager) (Manager, error)
}
