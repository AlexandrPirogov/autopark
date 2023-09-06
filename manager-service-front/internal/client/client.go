// package client hold URLs and interface ManagerHandler for requests
package client

// ApiGateway URL
const ApiGatewayHost = "http://api-gateway-nginx"

// Lists brandss
const AuthenticateURL = "/auth/login/manager"

const RegisterCarURL = "/autopark/car/register"
const ListCarsURL = "/autopark/car/list"

const RerfeshTokenCookieField = "refresh-token"

type Manager struct {
	Id              int    `json:"id"`
	EnterpriseTitle string `json:"e_title"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Login           string `json:"login"`
	Pwd             string `json:"pwd"`
}

type Enterprise struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Managers []Manager `json:"managers"`
}

type Car struct {
	UID   string `json:"uid"`
	Brand string `json:"brand"`
	Type  string `json:"type"`
}

type Brand struct {
	Brand string `json:"brand"`
}

type Client interface {
	// RegisterManager making request to register manager in auth-service
	//
	// Pre-cond: given client.Manager instance to register
	//
	// Post-cond: request was executed and result returned.
	// If request executes successfully returns Manager that was registeres and nil error
	// Otherwise returnes nil and error
	RegisterManager(m Manager) (Manager, error)

	RegisterEnterprise(e Enterprise) error
	ListEnterprises() ([]Enterprise, error)
	EnterpriseByID() (Enterprise, error)
}
