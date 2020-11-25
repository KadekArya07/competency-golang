package model

type LovCompetency struct {
	BaseModel
	KeyData string `json:"keyData" gorm:"unique"`
	ValData string `json:"valData"`
}

func (LovCompetency) TableName() string {
	return "cmp_lovs"
}
