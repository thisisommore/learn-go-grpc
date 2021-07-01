package main

import (
	client "github.com/thisisommore/calculator/calculator_client"
	server "github.com/thisisommore/calculator/calculator_server"
)

func main() {
	c := make(chan bool)
	go server.StartServer(c)

	<-c
	client.MakeRequest()
}
