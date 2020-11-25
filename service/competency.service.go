package service

import (
	"competency/config"
	"competency/dao"
	"competency/model"
	"competency/pojo"
	"log"

	"gorm.io/gorm"
)

type CompetencyService struct{}

var competencyDao = dao.CompetencyDao{}
var behaviourService = BehaviourService{}
var proficiencyService = ProficiencyService{}
var concernService = ConcernService{}
var trainingDetailService = TrainingDetailService{}
var nonTrainingService = NonTrainingService{}

func (CompetencyService) AddCompetency(pojoCompetency *pojo.PojoCompetency) (e error) {
	defer config.CatchError(&e)
	return g.Transaction(func(tx *gorm.DB) error {
		competency := pojoCompetency.Competency
		if err := competencyDao.AddCompetency(&competency, tx); err != nil {
			tx.Rollback()
			return err
		}

		if len(pojoCompetency.ListBehaviour) != 0 {
			for _, v := range pojoCompetency.ListBehaviour {
				behave := v.Behaviour
				behave.CompetencyID = competency.Id
				if err := behaviourService.AddBehaviour(&behave, tx); err != nil {
					tx.Rollback()
					return err
				}
				for _, f := range v.ListProficiency {
					f.BehaviourID = behave.Id
					if err := proficiencyService.AddProficiency(&f, tx); err != nil {
						tx.Rollback()
						return err
					}
				}
			}
		}

		pojoCompetency.Concern.CompetencyID = competency.Id
		if err := concernService.AddConcern(&pojoCompetency.Concern, tx); err != nil {
			log.Print(err)
			tx.Rollback()
			return err
		}

		if len(pojoCompetency.ListTraining) != 0 {
			for _, t := range pojoCompetency.ListTraining {
				trainingDetail := model.TrainingDetail{}
				trainingDetail.TrainingID = t.Id
				trainingDetail.CompetencyID = competency.Id
				if err := trainingDetailService.AddTrainingDetail(&trainingDetail, tx); err != nil {
					tx.Rollback()
					return err
				}
			}
		}

		pojoCompetency.NonTraining.CompetencyID = competency.Id
		if err := nonTrainingService.AddTraining(&pojoCompetency.NonTraining, tx); err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
}

func (CompetencyService) GetAllCompetency(page int, limit int) (pojoPagination pojo.PojoPagination, e error) {
	defer config.CatchError(&e)

	var listCompetency, err = competencyDao.GetCompetency(page, limit, "")
	var count, errC = competencyDao.GetCountCompetency(page, limit, "")
	if err != nil {
		return pojo.PojoPagination{}, err
	}
	if errC != nil {
		return pojo.PojoPagination{}, errC
	}

	pojoPagination.ListData = listCompetency
	pojoPagination.Count = count
	return pojoPagination, nil
}

func (CompetencyService) GetCompetencyBySearch(page int, limit int, inquiry string) (pojoPagination pojo.PojoPagination, e error) {
	defer config.CatchError(&e)

	var listCompetency, err = competencyDao.GetCompetency(page, limit, inquiry)
	var count, errC = competencyDao.GetCountCompetency(page, limit, inquiry)
	if err != nil {
		return pojo.PojoPagination{}, err
	}
	if errC != nil {
		return pojo.PojoPagination{}, errC
	}

	pojoPagination.ListData = listCompetency
	pojoPagination.Count = count
	return pojoPagination, nil
}

func (CompetencyService) GetCompetencyByJobId() (e error) {
	defer config.CatchError(&e)
	return nil
}
