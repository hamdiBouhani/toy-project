package gql

import (
	"bytes"
	"net/http"
	"strings"
	"toy-project/svc/configs"
	"toy-project/svc/gql/gqlutils"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/log"
	"github.com/pkg/errors"

	cors "github.com/rs/cors/wrapper/gin"
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
	// graphql API
	schemaString := GetRootSchemas()
	rootResolvers := &RootResolvers{
		Logger: s.Log,
	}

	schema, err := graphql.ParseSchema(
		schemaString,
		rootResolvers,
		graphql.Logger(&log.DefaultLogger{}),
		graphql.Tracer(&GQLTrace{Logger: s.Log}),
	)
	if err != nil {
		return errors.Wrap(err, "couldn't parse schema")
	}

	corsConfig := cors.Options{
		AllowedOrigins:   strings.Split(s.c.CORSHosts, ","),
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}

	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger())
	engine.Use(cors.New(corsConfig))

	authorized := engine.Group("/graphql")

	graphQLHandler := Handler{
		Schema: schema,
		Log:    s.Log,
	}

	authorized.Any("", graphQLHandler.Serve)

	openAccessed := engine.Group("/")
	{
		debugPage := bytes.Replace(gqlutils.GraphiQLPage, []byte("fetch('/'"), []byte("fetch('"+s.c.GqlDebugUrlPrefix+"'"), -1)
		openAccessed.GET("/graphql.html", func(c *gin.Context) {
			c.Data(http.StatusOK, "text/html; charset=utf-8", debugPage)
		})

		openAccessed.GET("/debug.html", func(c *gin.Context) {
			c.Data(http.StatusOK, "text/html; charset=utf-8", debugPage)
		})
	}

	httpServer := http.Server{
		Addr:    s.c.GqlAddress,
		Handler: engine,
	}

	// listening http server
	s.Log.Infoln("Starting http server listening on:", s.c.GqlAddress)
	if err := httpServer.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
