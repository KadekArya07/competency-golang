package dao

import (
	"competency/config"

	"gorm.io/gorm"
)

var g *gorm.DB

type BaseDao struct{}

func SetDao(gDB *gorm.DB) {
	g = gDB
}

func (BaseDao) Add(data interface{}) (e error) {
	defer config.CatchError(&e)
	result := g.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (BaseDao) AddTransaction(data interface{}, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	result := tx.Create(data)
	if result.Error == nil {
		return nil
	}
	return result.Error
}
