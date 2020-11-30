package service

import (
	"competency/config"
	"competency/dao"
	"competency/model"

	"gorm.io/gorm"
)

type NonTrainingService struct{}

var nonTrainingDao = dao.NonTrainingDao{}

func (NonTrainingService) AddTraining(data *model.NonTraining, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	return nonTrainingDao.AddNonTraining(data, tx)
}

func (NonTrainingService) GetNonTrainingByCompId(id string) (data []model.NonTraining, e error) {
	defer config.CatchError(&e)
	return nonTrainingDao.GetNonTrainingByCompId(id)
}
