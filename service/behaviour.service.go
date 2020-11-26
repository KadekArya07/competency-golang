package service

import (
	"competency/config"
	"competency/dao"
	"competency/model"

	"gorm.io/gorm"
)

type BehaviourService struct{}

var behaviourDao = dao.BehaviourDao{}

func (BehaviourService) AddBehaviour(behaviour *model.Behaviour, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	return behaviourDao.AddBehaviour(behaviour, tx)
}

func (BehaviourService) GetBehaviourByCompId(compId string) (listBehaviour []model.Behaviour, e error) {
	defer config.CatchError(&e)
	return behaviourDao.GetBehaviourByCompId(compId)
}
