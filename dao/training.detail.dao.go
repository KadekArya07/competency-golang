package dao

import (
	"competency/config"
	"competency/model"

	"gorm.io/gorm"
)

type TrainingDetailDao struct{}

func (TrainingDetailDao) AddTrainingDetail(data *model.TrainingDetail, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	result := tx.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}
