package model

type NonTraining struct {
	BaseModel
	CompetencyID string     `json:"competencyId" gorm:"column:comp_id"`
	Competency   Competency `json:"-" gorm:"foreignKey:CompetencyID"`
	NonTraining  string     `json:"non_training" gorm:"column:non_training"`
}

func (NonTraining) TableName() string {
	return "cmp_competencies_non_training"
}
