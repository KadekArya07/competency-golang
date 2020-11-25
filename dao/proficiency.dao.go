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
