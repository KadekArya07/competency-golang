package pojo

import "competency/model"

type PojoJobDetail struct {
	PojoBase
	ListCompetency []model.Competency
}
