package config

import (
	"context"
	"log"

	"google.golang.org/grpc"

	cr "competency/model"
	pb "competency/proto/model"
)

var Host = "192.168.15.98:1111"
var Ctx = context.Background()
var JobClient pb.JobServiceClient
var Empty = pb.Empty{}

var HostCredential = "camskoleksi.com:8091"
var CredentialClient cr.UserServiceClient

var HostEmployee = "192.168.15.65:8888"
var EmployeeClient pb.EmployeeServiceClient

func ConnectedJobService() {
	conn, err := grpc.Dial(Host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Not connected err =>", err)
	}

	JobClient = pb.NewJobServiceClient(conn)
}

func ConnectedCredential() {
	conn, err := grpc.Dial(HostCredential, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Not connected err =>", err)
	}

	CredentialClient = cr.NewUserServiceClient(conn)
}

func ConnectedEmployeeService() {
	conn, err := grpc.Dial(HostEmployee, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Not connected err =>", err)
	}

	EmployeeClient = pb.NewEmployeeServiceClient(conn)
}
