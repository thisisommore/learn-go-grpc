package main

import (
	client "calculator/calculator_client"
	server "calculator/calculator_server"
)

func main() {
	c := make(chan bool)
	go server.StartServer(c)

	<-c
	client.MakeRequest()
}
