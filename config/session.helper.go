package config

import "github.com/dgrijalva/jwt-go"

var users = make(map[string]interface{})

func SetSession(claims jwt.MapClaims) {
	defer CatchErrorGeneral()
	users = claims
}

func GetUser() (user *string) {
	userr := users["username"].(string)
	user = &userr
	return user
}
