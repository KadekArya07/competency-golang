package service

import (
	"competency/config"
	"competency/dao"
	"competency/model"
	"competency/pojo"
)

type TrainingService struct{}

var trainingDao = dao.TrainingDao{}

func (TrainingService) GetAllTraining() (listTraining []pojo.PojoLov, e error) {
	defer config.CatchError(&e)
	return trainingDao.GetAllTraining()
}

func (TrainingService) AddTraining(training *model.Training) (e error) {
	defer config.CatchError(&e)
	return baseDao.Add(training)
}
