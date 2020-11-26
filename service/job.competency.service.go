package service

import (
	"competency/config"
	"competency/dao"
	"competency/model"
)

type JobCompetencyService struct {
}

var jobCompetencyDao = dao.JobCompetencyDao{}

func (JobCompetencyService) GetByJobId(id string) (data []model.JobCompetency, e error) {
	defer config.CatchError(&e)
	return jobCompetencyDao.GetByJobId(id)
}

func (JobCompetencyService) DeleteJobById(jobId string) error {
	return jobCompetencyDao.DeleteByJobId(jobId)
}
