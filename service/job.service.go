package service

import (
	"competency/config"
	"competency/model"
	"competency/pojo"
	"encoding/json"

	"gorm.io/gorm"
)

type JobService struct{}

var jobCompetencyService = JobCompetencyService{}

func (JobService) GetAllJob() (data []*model.Jobs, e error) {
	defer config.CatchError(&e)
	res, err := config.JobClient.GetAllJob(config.Ctx, &config.Empty)
	if err != nil {
		panic("Not Connected !")
	}
	return res.Job, nil
}

func (JobService) GetJobById(id string) (data interface{}, e error) {
	defer config.CatchError(&e)
	res, err := config.JobClient.GetByIdJob(config.Ctx, &model.JobId{Id: id})
	if err != nil {
		panic(err)
	}
	job := pojo.PojoJobCompetency{}
	var maps map[string]interface{}
	apaaja, err := json.Marshal(res)
	json.Unmarshal(apaaja, &maps)
	job.Job = maps
	ress, err := jobCompetencyService.GetByJobId(job.Job["id"].(string))
	if err != nil {
		return nil, err
	}
	job.ListCompetency = ress
	return job, nil
}

func (JobService) AddCompetencyOnJob(data *pojo.PojoJobCompetency) (e error) {
	defer config.CatchError(&e)
	return g.Transaction(func(tx *gorm.DB) error {
		req := model.JobCompetency{}
		if len(data.ListCompetency) != 0 {
			for _, v := range data.ListCompetency {
				if v.BaseModel.Id == "" {
					req.CompetencyID = v.Competency.Id
					req.JobID = data.Job["id"].(string)

					if err := baseDao.AddTransaction(&req, tx); err != nil {
						tx.Rollback()
						return err
					}
				}
			}
		}
		return nil
	})

}
