package pojo

import "competency/model"

type PojoBehaviour struct {
	Behaviour       model.Behaviour     `json:"behave"`
	ListProficiency []model.Proficiency `json:"listProficiency"`
}
