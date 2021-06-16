package grpc

import (
	"log"
	"net"
	pb "toy-project/pb/toy-project"
	"toy-project/svc/configs"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	HostPort string
	Log      *logrus.Logger
	c        *configs.Config
}

func NewServer(logger *logrus.Logger, c *configs.Config) (*Server, error) {
	return &Server{
		HostPort: c.HostPort,
		Log:      logger,
		c:        c,
	}, nil
}

func (s *Server) Run() error {
	s.Log.Infoln("Starting Grpc API server listening on:", s.c.GRPCAddress)

	lis, err := net.Listen("tcp", s.c.GRPCAddress)
	if err != nil {
		log.Printf("Failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterToyProjectServer(grpcServer, NewAPI())

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
		return err
	}
	return nil
}
