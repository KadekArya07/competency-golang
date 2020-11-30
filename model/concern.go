package model

type Concern struct {
	BaseModel
	CompetencyID string     `json:"competencyId" gorm:"column:comp_id"`
	Competency   Competency `json:"-" gorm:"foreignKey:CompetencyID"`
	Concern      string     `json:"concern" gorm:"column:concern"`
}

func (Concern) TableName() string {
	return "cmp_competencies_concern"
}
