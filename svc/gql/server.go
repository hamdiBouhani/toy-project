package gql

import (
	"toy-project/svc/configs"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	*gin.Engine
	Log *logrus.Logger
	c   *configs.Config
}

func NewServer(logger *logrus.Logger, c *configs.Config) (*Server, error) {

	return &Server{
		Log: logger,
		c:   c,
	}, nil
}

func (s *Server) Run() error {
	return nil
}
