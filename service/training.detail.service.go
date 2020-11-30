package service

import (
	"competency/config"
	"competency/dao"
	"competency/model"

	"gorm.io/gorm"
)

type TrainingDetailService struct{}

var trainingDetailDao = dao.TrainingDetailDao{}

func (TrainingDetailService) AddTrainingDetail(data *model.TrainingDetail, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	return trainingDetailDao.AddTrainingDetail(data, tx)
}

// func ()
