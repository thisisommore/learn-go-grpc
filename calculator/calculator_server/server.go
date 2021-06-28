package server

import (
	pb "calculator/calculatorpb"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalulatorServiceServer
}

func (s *server) Add(c context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	addition := req.Num1 + req.Num2
	return &pb.AddResponse{Result: addition}, nil

}
func StartServer(c chan bool) {

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCalulatorServiceServer(grpcServer, &server{})
	c <- true
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Cannot serve: %v", err)
	}
}
