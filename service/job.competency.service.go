package service

import (
	"competency/config"
	"competency/dao"
	"competency/model"
	"context"
	"errors"
)

type JobCompetencyService struct {
	model.UnimplementedCompetencyServiceServer
}

var jobCompetencyDao = dao.JobCompetencyDao{}

func (JobCompetencyService) GetByJobId(id string) (data []model.JobCompetency, e error) {
	defer config.CatchError(&e)
	return jobCompetencyDao.GetByJobId(id)
}

func (JobCompetencyService) DeleteJobById(ctx context.Context, id *model.JobId) (*model.Response, error) {
	err := jobCompetencyDao.DeleteByJobId(id)
	if err != nil {
		return &model.Response{}, errors.New("Failed register")
	}
	return &model.Response{Code: "200", Message: "success"}, nil
}
