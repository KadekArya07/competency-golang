package dao

import (
	"competency/config"
	"competency/model"

	"gorm.io/gorm"
)

type ProficiencyDao struct{}

func (ProficiencyDao) AddProficiency(proficiency *model.Proficiency, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	result := tx.Create(proficiency)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (ProficiencyDao) GetProficiencyByBehaveId(id string) (data []model.Proficiency, e error) {
	defer config.CatchError(&e)
	var listProfi = []model.Proficiency{}
	result := g.Where("behave_id = ? ", id).Find(&listProfi)
	if result.Error == nil {
		return listProfi, nil
	}
	return data, result.Error
}
