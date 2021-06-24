# toy-project


it's Boilerplate of the GoLang Command-Line Applications.


That expose  two microservices (Rest & gRPC): 

* Available Commands

```
# => go run cmd/toy-project/main.go --help```

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
