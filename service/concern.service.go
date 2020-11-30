package service

import (
	"competency/config"
	"competency/dao"
	"competency/model"

	"gorm.io/gorm"
)

type ConcernService struct{}

var concernDao = dao.ConcernDao{}

func (ConcernService) AddConcern(data *model.Concern, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	return concernDao.AddConcern(data, tx)
}

func (ConcernService) GetConcernByCompId(id string) (data []model.Concern, e error) {
	defer config.CatchError(&e)
	return concernDao.GetConcernByCompId(id)
}
