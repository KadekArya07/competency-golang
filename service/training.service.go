package service

import (
	"competency/config"
	"competency/dao"
	"competency/pojo"
)

type TrainingService struct{}

var trainingDao = dao.TrainingDao{}

func (TrainingService) GetAllTraining() (listTraining []pojo.PojoLov, e error) {
	defer config.CatchError(&e)
	return trainingDao.GetAllTraining()
}
