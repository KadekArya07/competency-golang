package model

type Training struct {
	BaseModel
	Code string `json:"code"`
	Name string `json:"name"`
	Desc string `json:"desc" gorm:"column:description"`
}

func (Training) TableName() string {
	return "cmp_competencies_training"
}
