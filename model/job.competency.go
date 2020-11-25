package model

type JobCompetency struct {
	BaseModel
	JobID        string
	CompetencyID string     `json:"competencyId" gorm:"column:comp_id"`
	Competency   Competency `json:"competency" gorm:"foreignKey:CompetencyID"`
}

func (JobCompetency) TableName() string {
	return "cmp_competencies_job"
}
