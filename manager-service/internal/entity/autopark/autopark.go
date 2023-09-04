// package autopark holds structs of Brand and cars
// This can be used as a message for communication between services
package autopark

type Brand struct {
	Brand string `json:"brand"`
}

type Car struct {
	UID   string `json:"uid"`
	Brand string `json:"brand"`
	Type  string `json:"type"`
}
