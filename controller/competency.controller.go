package controller

import (
	"competency/config"
	"competency/pojo"
	"competency/service"

	"github.com/labstack/echo"
)

func SetCompetency(e *echo.Group) {
	e.POST("/competencies", addCompetency)
	e.GET("/competencies", getCompetency)
	e.GET("/competencies/search", getCompetencyBySearch)
	e.GET("/competencies/detail/:id", getCompetencyById)
}

var competencyService = service.CompetencyService{}

func addCompetency(c echo.Context) (e error) {
	defer config.CatchError(&e)
	competency := &pojo.PojoCompetency{}

	if err := c.Bind(competency); err != nil {
		return res(c, err)
	}

	if err := competencyService.AddCompetency(competency); err != nil {
		return resErr(c, err)
	}

	return resSuccess(c)
}

func getCompetency(c echo.Context) (e error) {
	defer config.CatchError(&e)
	list, err := competencyService.GetListCompetency(convInt(c.QueryParam("page")),
		convInt(c.QueryParam("limit")))
	if err != nil {
		return resErr(c, err)
	}
	return res(c, list)
}

func getCompetencyBySearch(c echo.Context) (e error) {
	defer config.CatchError(&e)
	list, err := competencyService.GetCompetencyBySearch(convInt(c.QueryParam("page")),
		convInt(c.QueryParam("limit")), c.QueryParam("inquiry"))
	if err != nil {
		return resErr(c, err)
	}
	return res(c, list)
}

func getCompetencyById(c echo.Context) (e error) {
	defer config.CatchError(&e)
	list, err := competencyService.GetById(c.Param("id"))
	if err != nil {
		return resErr(c, err)
	}
	return res(c, list)
}
