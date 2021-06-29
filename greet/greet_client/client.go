package client

import (
	"context"
	pb "greet/greetpb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func MakeRequest() {
	cc, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	s := pb.NewGreetServiceClient(cc)
	stream, err := s.GreetManyTimes(context.Background(), &pb.GreetManyTimesRequest{
		Name: "Om",
	})

	if err != nil {
		log.Fatalf("Failed to call GreetManyTime: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Print("End of response, thank you")
			break
		}
		log.Printf("Greetings received: %v", res.Greet)
	}
}
