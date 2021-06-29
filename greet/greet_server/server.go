package server

import (
	"fmt"
	pb "greet/greetpb"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetServiceServer
}

func (s *server) GreetManyTimes(req *pb.GreetManyTimesRequest, stream pb.GreetService_GreetManyTimesServer) error {
	for i := 0; i < 5; i++ {
		res := &pb.GreetManyTimesResponse{
			Greet: fmt.Sprintf("Hello %v", req.Name),
		}
		stream.Send(res)
		time.Sleep(2 * time.Second)
	}
	return nil
}

func StartServer(c chan bool) {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &server{})
	c <- true
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Cannot serve: %v", err)
	}
}
