package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/eabartlett/userservice/service"
)

func main() {
	user := new(userservice.UserService)
	err := rpc.Register(user)
	if err != nil {
		log.Fatalf("Format of userservice is not correct. %s", err)
	}
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatalf("error starting listening server. %s", e)
	}
	log.Println("Serving RPC handler")
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatalf("Error serving: %s", err)
	}
}
