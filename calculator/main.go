package main

import (
	client "github.com/thisisommore/learn-go-grpc/calculator/calculator_client"
	server "github.com/thisisommore/learn-go-grpc/calculator/calculator_server"
)

func main() {
	c := make(chan bool)
	go server.StartServer(c)

	<-c
	client.MakeRequest()
}
