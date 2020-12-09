package controller

import (
	"competency/config"
	"competency/model"

	"github.com/labstack/echo"
)

func SetTraining(e *echo.Group) {

	e.POST("/training", addTraining)
}

func addTraining(c echo.Context) (e error) {
	defer config.CatchError(&e)
	training := &model.Training{}

	if err := c.Bind(training); err != nil {
		return resErr(c, err)
	}

	if err := trainingService.AddTraining(training); err != nil {
		return resErr(c, err)
	}

	return resSuccess(c)
}
