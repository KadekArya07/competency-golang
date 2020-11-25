package model

type Proficiency struct {
	BaseModel
	BehaviourID string    `json:"behave_id" gorm:"column:behave_id"`
	Behaviour   Behaviour `json:"behaviour" gorm:"foreignKey:BehaviourID"`
	Rating      int       `json:"rating"`
	Desc        string    `json:"desc" gorm:"column:description"`
}

func (Proficiency) TableName() string {
	return "cmp_competencies_proficiency"
}
