package main;

import (
	"fmt"
	"golang.org/x/net/context"
	"mygrpc/service"
	"net"
	"log"
	"google.golang.org/grpc"
)

type myServer struct {
	service.UnimplementedMyServServer
}

func (s *myServer) SayName(ctx context.Context, req *service.Request) (*service.Response,error) {
	fmt.Printf("In sayName: req name is %v",req.Name)
	return &service.Response{GotName:"Ok I am admin GOOM and I got name"+req.Name},nil
}
func main()  {
	fmt.Println("Hii")
	lis,err := net.Listen("tcp",":5000")
	if err!=nil {
		log.Fatalf("Failed to listen: %v",err)
	}
	s := myServer{}
	grpcServer := grpc.NewServer()
	service.RegisterMyServServer(grpcServer,&s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v",err)
	}
}