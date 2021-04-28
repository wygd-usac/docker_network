package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/wygd-usac/gRPC/ServiceDefinition"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedCalculadoraServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Operar(ctx context.Context, in *pb.Operacion) (*pb.Resultado, error) {
	switch(in.Operador){
	case 1:
		return &pb.Resultado{Resultado: in.Op1 + in.Op2,Mensaje:"Suma realizada"}, nil
	default:
		return &pb.Resultado{Resultado: in.Op1 + in.Op2,Mensaje:"Se sumo por default"}, nil
		break;
	}
	return &pb.Resultado{Resultado: in.Op1 + in.Op2,Mensaje:"Se sumo por default"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculadoraServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}