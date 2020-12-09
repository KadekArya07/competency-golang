package service

import (
	"competency/config"
	"competency/dao"
	"competency/model"
	"competency/pojo"
	pb "competency/proto/model"
	"context"

	"gorm.io/gorm"
)

type CompetencyService struct {
}

var competencyDao = dao.CompetencyDao{}
var behaviourService = BehaviourService{}
var proficiencyService = ProficiencyService{}
var concernService = ConcernService{}
var trainingDetailService = TrainingDetailService{}
var nonTrainingService = NonTrainingService{}
var competencyServie = CompetencyService{}

func (CompetencyService) AddCompetency(pojoCompetency *pojo.PojoCompetency) (e error) {
	defer config.CatchError(&e)
	return g.Transaction(func(tx *gorm.DB) error {
		competency := pojoCompetency.Competency
		if err := baseDao.AddTransaction(&competency, tx); err != nil {
			tx.Rollback()
			return err
		}

		if len(pojoCompetency.ListBehaviour) != 0 {
			for _, v := range pojoCompetency.ListBehaviour {
				behave := v.Behaviour
				behave.CompetencyID = competency.Id
				if err := baseDao.AddTransaction(&behave, tx); err != nil {
					tx.Rollback()
					return err
				}
				for _, f := range v.ListProficiency {
					f.BehaviourID = behave.Id
					if err := baseDao.AddTransaction(&f, tx); err != nil {
						tx.Rollback()
						return err
					}
				}
			}
		}

		if len(pojoCompetency.ListBehaviour) != 0 {
			for _, concern := range pojoCompetency.ListConcern {
				concern.CompetencyID = competency.Id
				if err := baseDao.AddTransaction(&concern, tx); err != nil {
					tx.Rollback()
					return err
				}
			}
		}

		if len(pojoCompetency.ListTraining) != 0 {
			for _, t := range pojoCompetency.ListTraining {
				trainingDetail := model.TrainingDetail{}
				trainingDetail.TrainingID = t.Id
				trainingDetail.CompetencyID = competency.Id
				if err := baseDao.AddTransaction(&trainingDetail, tx); err != nil {
					tx.Rollback()
					return err
				}
			}
		}

		if len(pojoCompetency.ListNonTraining) != 0 {
			for _, non := range pojoCompetency.ListNonTraining {
				non.CompetencyID = competency.Id
				if err := baseDao.AddTransaction(&non, tx); err != nil {
					tx.Rollback()
					return err
				}
			}
		}
		return nil
	})
}

func (CompetencyService) GetListCompetency(page int, limit int) (pojoPagination pojo.PojoPagination, e error) {
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

func (CompetencyService) GetCompetencyById(id string) (competency model.Competency, e error) {
	defer config.CatchError(&e)
	return competencyDao.GetCompetencyById(id)
}

func (CompetencyService) GetAllCompetency(ctx context.Context, models *pb.Empty) (listCompetency *pb.CompetencyList, e error) {
	defer config.CatchError(&e)
	result, err := competencyDao.GetAllCompetency()
	listComp := []*pb.Competencies{}
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		listComp = append(listComp, &pb.Competencies{
			Id:   v.Id,
			Code: v.Code,
			Name: v.Name,
		})
	}
	listCompetency.ListCompetency = listComp
	return listCompetency, nil
}

func (CompetencyService) GetById(id string) (listData pojo.PojoCompetency, e error) {
	defer config.CatchError(&e)
	var pojoCompetency pojo.PojoCompetency
	result, err := competencyDao.GetById(id)
	if err == nil {
		if result.Id != "" {
			pojoCompetency.Competency = result
			return competencyServie.GetCompetencyDetail(pojoCompetency), nil
		}

	}
	return pojoCompetency, err
}

func (CompetencyService) GetCompetencyDetail(pojoCompetency pojo.PojoCompetency) (listData pojo.PojoCompetency) {
	id := pojoCompetency.Competency.Id

	result, err := behaviourService.GetBehaviourByCompId(id)
	if err != nil {
		panic(err)
	}

	for _, v := range result {
		data, err := proficiencyService.GetProficienyByBehaveId(v.Id)
		if err != nil {
			panic(err)
		}

		pojoCompetency.ListBehaviour = append(pojoCompetency.ListBehaviour,
			pojo.PojoBehaviour{
				Behaviour:       v,
				ListProficiency: data,
			})
	}

	listC, errs := concernService.GetConcernByCompId(id)
	if errs != nil {
		panic(errs)
	}
	pojoCompetency.ListConcern = listC

	listNon, errNon := nonTrainingService.GetNonTrainingByCompId(id)
	if errs != nil {
		panic(errNon)
	}
	pojoCompetency.ListNonTraining = listNon
	return pojoCompetency
}
