package main

import (
	"context"
	"log"

	"time"

	"google.golang.org/grpc"
	pb "github.com/wygd-usac/gRPC/ServiceDefinition"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculadoraServiceClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	r,err := c.Operar(ctx,&pb.Operacion{Op1:3,Op2:5,Operador:1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Resultado: %d", r.GetResultado())
}