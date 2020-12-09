package config

import (
	"competency/model"
	"errors"
	"fmt"
	"log"

	"github.com/labstack/echo"
)

func CatchError(e *error) {
	if err := recover(); err != nil {
		*e = fmt.Errorf("%v", err)
	}
}
func CatchErrorToken(token []string, e *error) error {
	if len(token) == 1 {
		return errors.New("token requ")
	}
	return nil
}

func CatchErrorGeneral() {
	if err := recover(); err != nil {
		log.Println("Error =>", err)
	}
}

func ValToken(c echo.Context) error {
	cs := c.(*model.CustomError)
	g := cs.CathError()
	if g != nil {
		panic(g)
	}
	return nil
}
