package main

import (
	"context"
	"log"
	"mygrpc/saynamepb"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer conn.Close()

	c := saynamepb.NewSayNameServiceClient(conn)

	res, err := c.SayName(context.Background(), &saynamepb.SayNameRequest{Name: "Om"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", res.GotName)
}
