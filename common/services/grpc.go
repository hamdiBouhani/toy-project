package services

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"toy-project/common/context"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc"
)

type GrpcService struct {
	context.DefaultService
	Server *grpc.Server
	port   string
}

const GRPC_SERVICE = "grpc_base"

//Id for the GRPC service
func (svc GrpcService) Id() string {
	return GRPC_SERVICE
}

//Configure GRPC service
func (svc *GrpcService) Configure(ctx *context.Context) error {
	port := flag.String("grpc_port", "8089", "Port for gRPC server to listen on")
	flag.Parse()

	svc.port = fmt.Sprintf(":%s", *port)
	svc.Server = grpc.NewServer()

	return svc.DefaultService.Configure(ctx)
}

//Start GRPC service and listen on port
func (svc *GrpcService) Start() error {
	return svc.Run()
}

func (svc *GrpcService) Run() error {
	lis, err := net.Listen("tcp", svc.port)
	if err != nil {
		return err
	}

	//Avoid blocking with gRPC service
	go func(svc *GrpcService) {
		if err = svc.Server.Serve(lis); err != nil {
			log.Panicf("Grpc Error: %s", err)
		}
	}(svc)

	log.Printf("Grpc listening on port: %s", svc.port)
	return nil
}

func (svc *GrpcService) RunBlocking() error {
	lis, err := net.Listen("tcp", svc.port)
	if err != nil {
		return err
	}

	log.Printf("Grpc listening on port: %s", svc.port)
	if err = svc.Server.Serve(lis); err != nil {
		log.Panicf("Grpc Error: %s", err)
	}

	return nil
}

//Shutdown the GRPC service
func (svc *GrpcService) Shutdown() {
	svc.Server.GracefulStop()
}

//UnmarshalToStruct to core DB struct for use within the service framework
// JSON conversion used as intermediary
func (svc *GrpcService) UnmarshalToStruct(msg proto.Message, out interface{}) error {
	m := jsonpb.Marshaler{
		OrigName:    true,
		EnumsAsInts: true,
	}
	jreq, err := m.MarshalToString(msg)
	if err != nil {
		log.Printf("Unable to create string from message")
		return err
	}

	return json.Unmarshal([]byte(jreq), out)
}

//StructToMessage convert to proto message via json
// JSON conversion used as intermediary
func (svc *GrpcService) StructToMessage(msg interface{}, out proto.Message) error {
	jreq, err := json.Marshal(msg)
	if err != nil {
		log.Printf("%s", jreq)
		return err
	}

	err = jsonpb.UnmarshalString(string(jreq), out)
	if err != nil {
		return err
	}

	return nil
}
