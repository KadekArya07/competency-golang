package model

type Behaviour struct {
	BaseModel
	CompetencyID string     `gorm:"column:comp_id" json:"compId"`
	Competency   Competency `gorm:"foreignKey:CompetencyID" json:"-"`
	Behaviour    string     `json:"behaviour"`
	Level        int        `json:"level"`
}

func (Behaviour) TableName() string {
	return "cmp_competencies_behaviour"
}
