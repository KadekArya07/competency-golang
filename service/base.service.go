package service

import (
	"competency/dao"

	"gorm.io/gorm"
)

var g *gorm.DB
var baseDao = dao.BaseDao{}

func SetService(gDB *gorm.DB) {
	g = gDB
}
