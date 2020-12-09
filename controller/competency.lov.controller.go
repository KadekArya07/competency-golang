package controller

import (
	"competency/model"

	"github.com/labstack/echo"
)

func SetLov(e *echo.Group) {
	e.POST("/lov", addLovCompetency)
	e.GET("/lov/training", getLovTraining)
	e.GET("/lov/compty", getCompetencyLov)
}

func addLovCompetency(c echo.Context) (e error) {
	defer catchError(&e)
	data := &model.LovCompetency{}

	if err := c.Bind(data); err != nil {
		return res(c, err)
	}

	if err := competencyLovService.AddLovCompetency(data); err != nil {
		return resErr(c, err)
	}

	return resSuccess(c)
}

func getLovTraining(c echo.Context) (e error) {
	defer catchError(&e)
	list, err := trainingService.GetAllTraining()
	if err != nil {
		return resErr(c, err)
	}
	return res(c, list)
}

func getCompetencyLov(c echo.Context) (e error) {
	defer catchError(&e)
	list, err := competencyLovService.GetAllCompetencyLov()
	if err != nil {
		return resErr(c, err)
	}
	return res(c, list)
}
