package controller

import (
	"competency/config"
	"competency/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

var competencyLovService = service.CompetencyLovService{}
var trainingService = service.TrainingService{}

func SetInit(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Rest API started")
	})
}

func res(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

func resErr(c echo.Context, err error) error {
	return c.String(http.StatusBadRequest, err.Error())
}

func resInvalid(c echo.Context, err error) error {
	return c.String(http.StatusUnauthorized, err.Error())
}

func catchError(e *error) {
	config.CatchError(e)
}

func resSuccess(c echo.Context) error {
	return c.String(http.StatusOK, "Success")
}

func convInt(val string) int {
	i, _ := strconv.Atoi(val)
	return i
}
