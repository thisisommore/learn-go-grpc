package main

import (
	client "greet/greet_client"
	server "greet/greet_server"
)

func main() {
	c := make(chan bool)
	go server.StartServer(c)
	<-c
	client.MakeRequest()
}
