package service

import (
	"competency/config"
	pb "competency/proto/model"
)

type EmployeeService struct {
}

func (EmployeeService) GetAllEmployee() (data *pb.Employees, e error) {
	defer config.CatchError(&e)

	result, err := config.EmployeeClient.GetEmployees(config.Ctx, &pb.Tokens{Token: config.ReqToken})

	if err != nil {
		panic("Proto Error ")
	}

	return result, nil
}
