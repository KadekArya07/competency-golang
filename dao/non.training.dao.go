package dao

import (
	"competency/config"
	"competency/model"

	"gorm.io/gorm"
)

type NonTrainingDao struct{}

func (NonTrainingDao) AddNonTraining(data *model.NonTraining, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	result := tx.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (NonTrainingDao) GetNonTrainingByCompId(id string) (data []model.NonTraining, e error) {
	defer config.CatchError(&e)
	var listNonTraining = []model.NonTraining{}
	result := g.Where("comp_id = ? ", id).Find(&listNonTraining)
	if result.Error == nil {
		return listNonTraining, nil
	}
	return data, result.Error
}
