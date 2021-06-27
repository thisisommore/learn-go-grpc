package main

import (
	"fmt"
	"log"
	"mygrpc/sayname/saynamepb"
	"net"

	"context"

	"google.golang.org/grpc"
)

type server struct {
	saynamepb.UnimplementedSayNameServiceServer
}

func (s *server) SayName(ctx context.Context, req *saynamepb.SayNameRequest) (*saynamepb.SayNameResponse, error) {
	fmt.Printf("In sayName: req name is %v", req.Name)
	return &saynamepb.SayNameResponse{GotName: "Ok I am admin GOOM and I got name " + req.Name}, nil
}
func main() {
	fmt.Println("Hii")
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := server{}
	grpcServer := grpc.NewServer()
	saynamepb.RegisterSayNameServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
