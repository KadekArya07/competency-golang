package service

import (
	"competency/config"
	"competency/dao"
	"competency/model"

	"gorm.io/gorm"
)

type ProficiencyService struct{}

var proficiencyDao = dao.ProficiencyDao{}

func (ProficiencyService) AddProficiency(proficiency *model.Proficiency, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	return proficiencyDao.AddProficiency(proficiency, tx)
}

func (ProficiencyService) GetProficienyByBehaveId(id string) (data []model.Proficiency, e error) {
	defer config.CatchError(&e)
	return proficiencyDao.GetProficiencyByBehaveId(id)
}
