package server

import (
	"competency/model"
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

	model.RegisterCompetencyServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve with err =>", err)
	}

}

type server struct {
	model.UnimplementedCompetencyServiceServer
}

var jobCompetencyService = service.JobCompetencyService{}

func (*server) DeleteJobById(ctx context.Context, id *model.JobId) (*model.Response, error) {
	result, err := jobCompetencyService.GetByJobId(id.Id)
	log.Print(err)
	if err != nil {
		return &model.Response{Code: "400"}, err
	}

	if len(result) != 0 {
		vs := model.JobCompetency{}
		for _, vs = range result {
			job := &model.JobId{Id: vs.Id}
			errs := jobCompetencyService.DeleteJobById(job.Id)
			if errs != nil {
				return &model.Response{Code: "400", Message: "Failed delete !"}, err
			}
		}
	}

	return &model.Response{Code: "200", Message: "success"}, nil
}
