
proto:
	protoc -Iproto --go_out=plugins=grpc:pb proto/toy-project/toy-project.proto

startapp:
	go run cmd/toy-project/main.go serve
