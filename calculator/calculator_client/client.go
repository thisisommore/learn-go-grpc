package calculator_client

import (
	"context"
	"log"

	pb "github.com/thisisommore/learn-go-grpc/calculator/calculatorpb"

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

	numbers := []int32{10, 9, 3}
	stream, err := c.Sum(context.Background())
	if err != nil {
		log.Fatalf("Cannot call Sum: %v", err)
	}
	for _, number := range numbers {
		stream.Send(&pb.SumRequest{
			Num: number,
		})
	}
	sumRes, _ := stream.CloseAndRecv()
	log.Printf("Sum is: %v", sumRes.Result)
}
