package pojo

import "competency/model"

type PojoCompetency struct {
	Competency      model.Competency    `json:"competency"`
	ListBehaviour   []PojoBehaviour     `json:"listBehave"`
	ListConcern     []model.Concern     `json:"listConcern"`
	ListTraining    []model.Training    `json:"listTraining"`
	ListNonTraining []model.NonTraining `json:"listNonTraining"`
}
