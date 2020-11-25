package controller

import (
	"competency/pojo"
	"competency/service"

	"github.com/labstack/echo"
)

var jobService = service.JobService{}

func SetJob(e *echo.Group) {
	e.GET("/job/getall", getAllJob)
	e.GET("/job/details/:id", getJobById)
	e.POST("/job/add-competency", addCompetencyOnJob)
}

func getAllJob(c echo.Context) (e error) {
	defer catchError(&e)
	ress, err := jobService.GetAllJob()
	if err != nil {
		return resErr(c, err)
	}
	return res(c, ress)
}

func getJobById(c echo.Context) (e error) {
	defer catchError(&e)
	ress, err := jobService.GetJobById(c.Param("id"))
	if err != nil {
		return resErr(c, err)
	}
	return res(c, ress)
}

func addCompetencyOnJob(c echo.Context) (e error) {
	defer catchError(&e)
	data := &pojo.PojoJobCompetency{}
	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	if err := jobService.AddCompetencyOnJob(data); err != nil {
		return resErr(c, err)
	}
	return resSuccess(c)
}
