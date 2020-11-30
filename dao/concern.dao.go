package dao

import (
	"competency/config"
	"competency/model"
	"log"

	"gorm.io/gorm"
)

type ConcernDao struct{}

func (ConcernDao) AddConcern(data *model.Concern, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	result := tx.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (ConcernDao) GetConcernByCompId(id string) (data []model.Concern, e error) {
	defer config.CatchError(&e)
	var listConcern = []model.Concern{}
	result := g.Where("comp_id = ?", id).Find(&listConcern)
	log.Print(result)
	if result.Error == nil {
		return listConcern, nil
	}
	return data, result.Error
}
