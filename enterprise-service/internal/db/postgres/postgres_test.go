package postgres

import "testing"

// Use it to test connection to your database
func TestTryConnect(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("error while connection to db %v", r)
		}
	}()
	GetInstance()
}
