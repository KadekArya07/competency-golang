package dao

import (
	"competency/config"
	"competency/model"

	"gorm.io/gorm"
)

type BehaviourDao struct{}

func (BehaviourDao) AddBehaviour(behaviour *model.Behaviour, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	result := tx.Create(behaviour)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (BehaviourDao) GetBehaviourByCompId(compId string) (listBehaviour []model.Behaviour, e error) {
	defer config.CatchError(&e)
	listBehaviour = []model.Behaviour{}
	result := g.Where("comp_id = ? ", compId).Find(&listBehaviour)
	if result.Error == nil {
		return listBehaviour, nil
	}
	return listBehaviour, result.Error
}
