package pojo

import "competency/model"

type PojoCompetency struct {
	Competency    model.Competency  `json:"competency"`
	ListBehaviour []PojoBehaviour   `json:"listBehave"`
	Concern       model.Concern     `json:"concern"`
	ListTraining  []model.Training  `json:"listTraining"`
	NonTraining   model.NonTraining `json:"nonTraining"`
}
