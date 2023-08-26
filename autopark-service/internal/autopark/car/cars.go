// package car holds structs for Brand and Car
package car

type Brand struct {
	Brand string `json:"brand"`
}

type Car struct {
	UID   string `json:"uid"`
	Brand string `json:"brand"`
	Type  string `json:"type"`
}
