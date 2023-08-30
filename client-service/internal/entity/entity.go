package entity

type ClientCreds struct {
	UserID int `json:"u_id"`
}

type Booking struct {
	ID     int `json:"id"`
	UserID int `json:"u_id"`
	CarID  int `json:"c_id"`
}
