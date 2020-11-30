package controller

import (
	"competency/config"
	"competency/service"

	"github.com/labstack/echo"
)

func SetEmployee(e *echo.Group) {
	e.GET("/employees", getAllEmployees)
}

var employeeService = service.EmployeeService{}

func getAllEmployees(c echo.Context) (e error) {

	defer config.CatchError(&e)
	result, err := employeeService.GetAllEmployee()
	if err != nil {
		return resErr(c, err)
	}
	return res(c, result)
}
