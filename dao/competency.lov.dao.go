package dao

import (
	"competency/config"
	"competency/pojo"
	"strings"
)

type CompetencyLovDao struct{}

func (CompetencyLovDao) GetAllCompetencyLov() (listCompetencyLov []pojo.PojoLov, e error) {
	defer config.CatchError(&e)
	var sb strings.Builder
	sb.WriteString("SELECT id,key_data,val_data ")
	sb.WriteString("FROM cmp_lovs ")

	rows, err := g.Raw(sb.String()).Rows()

	for rows.Next() {
		listCompetencyLov = []pojo.PojoLov{}
		pojoLov := pojo.PojoLov{}
		var key string
		var val string

		rows.Scan(&pojoLov.Key, &key, &val)
		pojoLov.Value = key + " - " + val
		listCompetencyLov = append(listCompetencyLov, pojoLov)
	}
	defer rows.Close()
	return listCompetencyLov, err
}
