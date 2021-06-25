package main

import (
	"log"
	"context"	
	"google.golang.org/grpc"
	"mygrpc/service"
)

func main()  {
	var conn *grpc.ClientConn
	conn,err := grpc.Dial(":5000",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error: %v",err)
	}

	defer conn.Close()

	c := service.NewMyServClient(conn)

	res,err := c.SayName(context.Background(),&service.Request{Name:"Om"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", res.GotName)
}