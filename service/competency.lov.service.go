package service

import (
	"competency/config"
	"competency/dao"
	"competency/model"
	"competency/pojo"
	pb "competency/proto/model"
)

type CompetencyLovService struct {
	pb.UnimplementedCompetencyServiceServer
}

var comptencyLovDao = dao.CompetencyLovDao{}

func (CompetencyLovService) AddLovCompetency(lov *model.LovCompetency) (e error) {
	defer config.CatchError(&e)
	return baseDao.Add(lov)
}

func (CompetencyLovService) GetAllCompetencyLov() (listCompetencyLov []pojo.PojoLov, e error) {
	defer config.CatchError(&e)
	return comptencyLovDao.GetAllCompetencyLov()
}
