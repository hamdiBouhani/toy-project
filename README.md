# toy-project


it's Boilerplate of the GoLang Command-Line Applications.


That expose  two microservices (Rest & gRPC): 

* Available Commands

```
# => go run cmd/toy-project/main.go --help

Usage:
  toy-project [flags]
  toy-project [command]

Available Commands:
  help        Help about any command
  serve       Connect to the storage and begin serving requests.

Flags:
  -h, --help      help for toy-project
  -v, --verbose   verbose output

Use "toy-project [command] --help" for more information about a command.

```

* Serve CMD flags

```
# => go run cmd/toy-project/main.go serve --help

Connect to the storage and begin serving requests.

Usage:
  toy-project serve [flags]

Flags:
  -a, --address string        address to listen on (default ":8080")
      --cors-hosts string     cors hosts, separated by comma (default "*")
      --dsn string            db url (default "postgres://postgres@localhost:5432/toy_project_db?sslmode=disable")
      --grpc-address string   GRPC address of configurator (default ":9301")
  -h, --help                  help for serve

Global Flags:
  -v, --verbose   verbose output
app starded successfully !!
```

Also, there is Make cmd to start project:
> make -B start-app

# Project structure

```
.
├── cmd
│   └── toy-project
│       └── main.go
├── go.mod
├── go.sum
├── Makefile
├── pb
│   └── toy-project
│       └── toy-project.pb.go
├── proto
│   └── toy-project
│       └── toy-project.proto
├── README.md
├── storage
│   ├── mongo
│   │   └── mongo.go
│   ├── pg
│   │   ├── connect.go
│   │   ├── migrate.go
│   │   └── pg.go
│   └── storage.go
├── svc
│   ├── cmd
│   │   └── serve
│   │       └── serve.go
│   ├── configs
│   │   └── configs.go
│   ├── gql
│   │   └── server.go
│   ├── grpc
│   │   ├── api.go
│   │   └── server.go
│   ├── rest
│   │   └── server.go
│   └── server
│       └── server.go
└── thirdparty
    └── tokenverifier.go

18 directories, 20 files
```
