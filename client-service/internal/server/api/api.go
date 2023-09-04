package api

import (
	"fmt"
	"net/http"
	"strings"
)

func RetrieveRefreshToken(r *http.Request) (string, error) {
	var cookie *http.Cookie
	cookie, err := r.Cookie("refresh-token")

	if cookie == nil || err != nil {
		return "", fmt.Errorf("error while reading token %v", err)
	}

	tokenVal := cookie.Value[strings.Index(cookie.Value, "=")+1:]

	return tokenVal, err
}
