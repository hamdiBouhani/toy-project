package gql

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/sirupsen/logrus"
)

// Handler a graphql Handle responds to an HTTP request.
type Handler struct {
	Schema               *graphql.Schema
	Log                  logrus.FieldLogger
	EnableDashboardCache bool
}
