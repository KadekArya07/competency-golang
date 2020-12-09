package config

import (
	"competency/model"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

const secret = "secret"

var ReqToken string

func MiddlewareCredential(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (e error) {
		ReqToken = c.Request().Header.Get("Authorization")
		splitToken := strings.Split(ReqToken, "Bearer ")
		es := CatchErrorToken(splitToken, &e)
		if es != nil {
			return c.String(http.StatusUnauthorized, "Token is Required")

		} else {
			ReqToken = splitToken[1]
			// err := CheckCredentialToken(ReqToken)
			// if err != nil {
			// 	return c.JSON(http.StatusUnauthorized, err.Error())
			// }

			ss := &model.CustomError{c, nil}
			GetClaims(ReqToken)
			return next(ss)
		}

	}
}

func CheckCredentialToken(token string) error {
	res, err := CredentialClient.ValidateToken(Ctx, &model.Token{Data: token})

	if err != nil {
		desc := strings.Split(err.Error(), "desc = ")
		err = errors.New(desc[1])
		log.Println("Error validate =>", err)
		return err
	}

	log.Println("Success validate =>", res)
	return nil
}

func GetClaims(tokenStr string) {
	defer CatchErrorGeneral()

	token, _ := jwt.Parse(tokenStr, nil)
	claims, _ := token.Claims.(jwt.MapClaims)

	SetSession(claims)
}
