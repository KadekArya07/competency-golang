package dao

import (
	"competency/config"
	"competency/model"
)

type JobCompetencyDao struct{}

func (JobCompetencyDao) GetByJobId(id string) (data []model.JobCompetency, e error) {
	defer config.CatchError(&e)
	var listJobCompetency []model.JobCompetency
	result := g.Where("job_id", id).Preload("Competency").Find(&listJobCompetency)
	if result.Error == nil {
		return listJobCompetency, nil
	}
	return nil, result.Error
}

func (JobCompetencyDao) DeleteByJobId(id string) (e error) {
	defer config.CatchError(&e)
	result := g.Where("id = ?", id).Delete(model.JobCompetency{})
	if result.Error == nil {
		return nil
	}

	return result.Error
}
