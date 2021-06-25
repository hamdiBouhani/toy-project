package server

import (
	"toy-project/svc/configs"
	"toy-project/svc/gql"
	grpcServer "toy-project/svc/grpc"
	"toy-project/svc/rest"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger     logrus.FieldLogger
	restServer *rest.Server
	grpcServer *grpcServer.Server
	gqlServer  *gql.Server
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

	gqlServer, err := gql.NewServer(logger, c)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a new rest server")
	}

	return &Server{
		logger:     logger,
		restServer: restServer,
		grpcServer: grpcServer,
		gqlServer:  gqlServer,
	}, nil
}

func (s *Server) Run() error {

	if s.grpcServer != nil {
		go func() { s.grpcServer.Run() }()
	} else {
		s.logger.Infoln("no grpc server started")
	}

	if s.gqlServer != nil {
		go func() { s.gqlServer.Run() }()
	} else {
		s.logger.Infoln("no graphql server started")
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
