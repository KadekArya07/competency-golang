package server

import (
	"competency/config"
	"competency/model"
	pb "competency/proto/model"
	"competency/service"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

func SetProto() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("Failed to listen with err =>", err)
	}

	s := grpc.NewServer()

	pb.RegisterCompetencyServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve with err =>", err)
	}

}

type server struct {
	pb.UnimplementedCompetencyServiceServer
}

var jobCompetencyService = service.JobCompetencyService{}
var competencyService = service.CompetencyService{}

func (*server) DeleteJobById(ctx context.Context, id *pb.JobId) (*pb.Response, error) {
	result, err := jobCompetencyService.GetByJobId(id.Id)
	log.Print(err)
	if err != nil {
		return &pb.Response{Code: "400"}, err
	}

	if len(result) != 0 {
		vs := model.JobCompetency{}
		for _, vs = range result {
			job := &pb.JobId{Id: vs.Id}
			errs := jobCompetencyService.DeleteJobById(job.Id)
			if errs != nil {
				return &pb.Response{Code: "400", Message: "Failed delete !"}, err
			}
		}
	}

	return &pb.Response{Code: "200", Message: "success"}, nil
}

func (*server) GetCompetencyById(ctx context.Context, models *pb.CompId) (competency *pb.Competencies, e error) {
	defer config.CatchError(&e)
	result, err := competencyService.GetCompetencyById(models.Id)
	if err != nil {
		return competency, err
	}
	return &pb.Competencies{
		Id:   *&result.Id,
		Code: *&result.Code,
		Name: *&result.Name,
	}, nil
}

func (*server) GetCompetencyByJobId(ctx context.Context, jobId *pb.CompId) (jobs *pb.CompetencyList, e error) {
	defer config.CatchError(&e)
	result, err := jobCompetencyService.GetCompetencyByJobId(jobId.Id)
	if err != nil {
		return jobs, err
	}

	jobss := &pb.CompetencyList{}
	for _, v := range result {
		jobss.ListCompetency = append(jobss.ListCompetency, &pb.Competencies{
			Id:   v.Competency.Id,
			Code: v.Competency.Code,
			Name: v.Competency.Name,
		})
	}
	return jobss, err
}
