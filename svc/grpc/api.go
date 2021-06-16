package grpc

import (
	"context"
	pb "toy-project/pb/toy-project"
)

type api struct {
}

func NewAPI() pb.ToyProjectServer {
	return &api{}
}

func (a *api) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingReply, error) {
	return &pb.PingReply{Up: true}, nil
}
