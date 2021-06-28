package client

import (
	pb "calculator/calculatorpb"
	"context"
	"log"

	"google.golang.org/grpc"
)

func MakeRequest() {
	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalulatorServiceClient(conn)
	res, err := c.Add(context.Background(), &pb.AddRequest{Num1: 10, Num2: 21})
	if err != nil {
		log.Fatalf("Cannot call Add: %v", err)
	}
	log.Printf("Got res: %v", res)

}
