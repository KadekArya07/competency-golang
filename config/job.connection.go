package config

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "competency/model"
)

var Host = "192.168.15.98:1111"
var Ctx = context.Background()
var JobClient pb.JobServiceClient
var Empty = pb.Empty{}

func Connected() {
	conn, err := grpc.Dial(Host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Not connected err =>", err)
	}

	JobClient = pb.NewJobServiceClient(conn)
}
