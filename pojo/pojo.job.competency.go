package pojo

import "competency/model"

type PojoJobCompetency struct {
	Job            map[string]interface{} `json:"job"`
	ListCompetency []model.JobCompetency  `json:"listCompetency"`
}
