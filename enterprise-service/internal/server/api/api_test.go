package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterEntityCode(t *testing.T) {
	expected := http.StatusCreated
	request := httptest.NewRequest(http.MethodPost, "/register", nil)

	w := httptest.NewRecorder()
	h := http.HandlerFunc(RegisterEnterprise)
	h.ServeHTTP(w, request)
	res := w.Result()

	defer res.Body.Close()

	assert.Equal(t, expected, res.StatusCode)
}

func TestReadEntityCode(t *testing.T) {
	expected := http.StatusOK
	request := httptest.NewRequest(http.MethodGet, "/read", nil)

	w := httptest.NewRecorder()
	h := http.HandlerFunc(ReadEnerprises)
	h.ServeHTTP(w, request)
	res := w.Result()

	defer res.Body.Close()

	assert.Equal(t, expected, res.StatusCode)
}

func TestDeleteEntityCode(t *testing.T) {
	expected := http.StatusNoContent
	request := httptest.NewRequest(http.MethodPost, "/delete", nil)

	w := httptest.NewRecorder()
	h := http.HandlerFunc(DeleteEnterprise)
	h.ServeHTTP(w, request)
	res := w.Result()

	defer res.Body.Close()

	assert.Equal(t, expected, res.StatusCode)
}

func TestUpdateEntityCode(t *testing.T) {
	expected := http.StatusOK
	request := httptest.NewRequest(http.MethodPost, "/update", nil)

	w := httptest.NewRecorder()
	h := http.HandlerFunc(UpdateEnterprises)
	h.ServeHTTP(w, request)
	res := w.Result()

	defer res.Body.Close()

	assert.Equal(t, expected, res.StatusCode)
}
