package dao

import (
	"competency/config"
	"competency/pojo"
	"strings"
)

type TrainingDao struct{}


func (TrainingDao) GetAllTraining() (data []pojo.PojoLov, e error) {
	defer config.CatchError(&e)
	var sb strings.Builder
	sb.WriteString("SELECT id,code,name ")
	sb.WriteString("FROM cmp_competencies_training ")

	rows, err := g.Raw(sb.String()).Rows()
	listTraining := []pojo.PojoLov{}
	for rows.Next() {

		pojoLov := pojo.PojoLov{}
		var code string
		var name string

		rows.Scan(&pojoLov.Key, &code, &name)
		pojoLov.Value = code + " - " + name
		listTraining = append(listTraining, pojoLov)
	}
	defer rows.Close()
	return listTraining, err
}
