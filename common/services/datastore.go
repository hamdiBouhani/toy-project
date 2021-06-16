package services

import (
	"flag"
	"toy-project/common/context"
)

type Datastore struct {
	context.DefaultService

	username string
	password string
	database string
	host     string
}

const DATASTORE_SVC = "datastore"

var dbUser = flag.String("db_username", "root", "Database username")
var dbPass = flag.String("db_password", "", "Database username")
var dbHost = flag.String("db_host", "localhost", "Database username")
var dbDatabase = flag.String("db_database", "loading", "Database username")

//Configures the connection params based on the flags provided
func (ds *Datastore) Configure(ctx *context.Context) error {
	flag.Parse()

	ds.username = *dbUser
	ds.password = *dbPass
	ds.host = *dbHost
	ds.database = *dbDatabase

	return ds.DefaultService.Configure(ctx)
}
