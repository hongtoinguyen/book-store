package SetUp

import (
	"book-store/model/Jwt"
	"net/http"
)

var ErrorMap map[string]int //global variables

//InitStuffs ...
func InitStuffs() {
	Jwt.LoadRSAKeys()
	ErrorMap = map[string]int{
		"":                              http.StatusServiceUnavailable,
		"Email or password is invalid":  http.StatusForbidden,
		"You have not registered yet":   http.StatusNotAcceptable,
		"Password is not right":         http.StatusNotAcceptable,
		"You have registered":           http.StatusNotAcceptable,
		"Can't hash password":           http.StatusServiceUnavailable,
		"Can't decode hash of password": http.StatusServiceUnavailable,
		"Token is invalid":              http.StatusUnauthorized,
	}
}
