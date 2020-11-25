package model

type TrainingDetail struct {
	BaseModel
	TrainingID   string     `json:"trainingId" gorm:"column:training_id"`
	Training     Training   `json:"training" gorm:"foreignKey:TrainingID"`
	CompetencyID string     `json:"competencyId" gorm:"column:competency_id"`
	Competency   Competency `json:"competency" gorm:"foreignKey:CompetencyID"`
}

func (TrainingDetail) TableName() string {
	return "cmp_competencies_training_detail"
}
