package config

import (
	"competency/model"
	"errors"
	"log"
	"net/http"
	"strings"

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
			err := CheckCredentialToken(ReqToken)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, err.Error())
			}

			ss := &model.CustomError{c, nil}
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

// var ReqToken string
// var BaseUrlEmployee = "http://camskoleksi.com:8091"

// func MiddlewareCredential(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) (e error) {
// 		ReqToken = c.Request().Header.Get("Authorization")
// 		splitToken := strings.Split(ReqToken, "Bearer ")
// 		es := config.CatchErrorToken(splitToken, &e)
// 		if es != nil {
// 			return c.String(http.StatusUnauthorized, "Token is Required")

// 		} else {
// 			ReqToken = splitToken[1]
// 			err := CheckCredentialToken(ReqToken)
// 			if err != nil {
// 				return c.JSON(http.StatusUnauthorized, err.Error())
// 			}

// 			ss := &model.CustomError{c, nil}
// 			return next(ss)
// 		}

// 	}
// }

// func CheckCredentialToken(token string) error {
// 	res, err := config.Client.ValidateToken(config.Ctx,
// 		&pb.Token{Data: token})

// 	if err != nil {
// 		desc := strings.Split(err.Error(), "desc = ")
// 		err = errors.New(desc[1])
// 		log.Println("Error validate =>", err)
// 		return err
// 	}

// 	log.Println("Success validate =>", res)
// 	return nil
// }

// func GetRequest(uri string) (bodyResp []byte, e error) {
// 	var bearer = "Bearer " + ReqToken
// 	defer config.CatchError(&e)
// 	req, err := http.NewRequest("GET", BaseUrlEmployee+uri, nil)
// 	req.Header.Add("Authorization", bearer)
// 	req.Header.Add("Content-Type", "application/json;charset=UTF-8")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Println("Error on response =>", err)
// 	}
// 	defer resp.Body.Close()

// 	return ioutil.ReadAll(resp.Body)
// }
// }
