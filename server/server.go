package server

import (
	"competency/model"
	"competency/service"
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
	model.RegisterCompetencyServiceServer(s, &service.JobCompetencyService{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve with err =>", err)
	}
}
