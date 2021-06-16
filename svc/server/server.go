package server

import (
	"toy-project/svc/configs"
	grpcServer "toy-project/svc/grpc"
	"toy-project/svc/rest"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger     logrus.FieldLogger
	restServer *rest.Server
	grpcServer *grpcServer.Server
}

func NewServer(c *configs.Config) (*Server, error) {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	restServer, err := rest.NewServer(logger, c)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a new rest server")
	}

	grpcServer, err := grpcServer.NewServer(logger, c)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a new rest server")
	}

	return &Server{
		logger:     logger,
		restServer: restServer,
		grpcServer: grpcServer,
	}, nil
}

func (s *Server) Run() error {

	if s.grpcServer != nil {
		go func() { s.grpcServer.Run() }()
	} else {
		s.logger.Infoln("no grpc server started")
	}

	if s.restServer != nil {
		err := s.restServer.Run()
		if err != nil {
			return err
		}
	} else {
		s.logger.Infoln("no rest server started")
	}

	return nil
}
