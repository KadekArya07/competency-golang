package model

type Competency struct {
	BaseModel
	LovCompetencyID string        `json:"lov_cmp_id" gorm:"column:lov_comp"`
	LovCompetency   LovCompetency `json:"lov_cmp" gorm:"foreignKey:LovCompetencyID"`
	Code            string        `json:"code" gorm:"column:code;unique"`
	Name            string        `json:"name" gorm:"name"`
	Desc            string        `json:"desc" gorm:"description"`
	MaxLevel        int           `json:"max_level" gorm:"max_level"`
}

func (Competency) TableName() string {
	return "cmp_competencies"
}
