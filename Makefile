
proto:
	protoc -Iproto --go_out=plugins=grpc:pb proto/toy-project/toy-project.proto

start-app:
	go run cmd/toy-project/main.go serve
