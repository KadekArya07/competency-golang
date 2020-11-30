package dao

import (
	"competency/config"
	"competency/model"
	"log"
	"strings"

	"gorm.io/gorm"
)

type CompetencyDao struct{}

func (CompetencyDao) GetById(id string) (data model.Competency, e error) {
	defer config.CatchError(&e)
	var competency model.Competency
	result := g.Where("id = ? ", id).Preload("LovCompetency").Find(&competency)
	if result.Error == nil {
		return competency, nil
	}
	return data, result.Error
}
func (CompetencyDao) AddCompetency(competency *model.Competency, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	result := tx.Create(competency)
	if result.Error == nil {
		return nil
	}
	return result.Error
}

func (CompetencyDao) GetCompetency(page int, limit int, inquiry string) (listCompetency []map[string]interface{}, e error) {
	defer config.CatchError(&e)
	var sb strings.Builder
	sb.WriteString("SELECT c.id,c.code,c.name,c.desc ")
	sb.WriteString("FROM ( ")
	sb.WriteString("SELECT cc.id,cc.code,cc.name,cc.desc ")
	sb.WriteString("FROM cmp_competencies cc ")
	sb.WriteString("LIMIT ? OFFSET ? ) as c ")
	sb.WriteString("WHERE 1=1 ")

	if inquiry != "" {
		sb.WriteString(" AND ( POSITION(LOWER('")
		sb.WriteString(inquiry)
		sb.WriteString("') in LOWER(CONCAT(")
		sb.WriteString("c.id,c.code,c.name,c.desc")
		sb.WriteString("))) > 0 )")
	}

	rows, err := g.Raw(sb.String(), limit, (page-1)*limit).Rows()
	listCompetency = make([]map[string]interface{}, 0, 0)

	for rows.Next() {
		var id string
		var code string
		var name string
		var desc string

		competency := make(map[string]interface{})
		rows.Scan(&id, &code, &name, &desc)

		competency["id"] = id
		competency["code"] = code
		competency["name"] = name
		competency["desc"] = desc

		listCompetency = append(listCompetency, competency)
	}
	defer rows.Close()
	return listCompetency, err
}

func (CompetencyDao) GetCountCompetency(page int, limit int, inquiry string) (count int, e error) {
	defer config.CatchError(&e)
	var sb strings.Builder
	sb.WriteString("SELECT COUNT(*)")
	sb.WriteString("FROM ( ")
	sb.WriteString("SELECT cc.id,cc.code,cc.name,cc.desc ")
	sb.WriteString("FROM cmp_competencies cc ")
	sb.WriteString("LIMIT ? OFFSET ? ) as c ")
	sb.WriteString("WHERE 1=1 ")

	if inquiry != "" {
		sb.WriteString(" AND ( POSITION(LOWER('")
		sb.WriteString(inquiry)
		sb.WriteString("') in LOWER(CONCAT(")
		sb.WriteString("c.id,c.code,c.name,c.desc")
		sb.WriteString("))) > 0 )")
	}

	rows, err := g.Raw(sb.String(), limit, (page-1)*limit).Rows()

	for rows.Next() {
		rows.Scan(&count)
	}
	defer rows.Close()

	return count, err
}

func (CompetencyDao) GetCompetencyById(id string) (competency model.Competency, e error) {
	defer config.CatchError(&e)
	var competencies = model.Competency{}
	result := g.Where("id = ?", id).Find(&competencies)
	log.Print(result)
	if result.Error == nil {
		return competencies, nil
	}

	return competencies, result.Error
}

func (CompetencyDao) GetAllCompetency() (listCompetency []model.Competency, e error) {
	defer config.CatchError(&e)
	var competencies = []model.Competency{}
	result := g.Find(&competencies)
	if result.Error == nil {
		return competencies, nil
	}

	return competencies, result.Error
}
