package main

import (
	"log"
	"net/rpc"

	"github.com/eabartlett/userservice/service"
)

func main() {
	// connect to client
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatalf("Unable to connect to service. Error %s", err)
	}
	// create arguments
	args := &userservice.GetByIdArgs{12345}
	// object for storing result
	var result userservice.Result
	//make call
	err = client.Call("UserService.GetUserByID", args, &result)
	if err != nil {
		log.Fatalf("Error calling user service. Error %s", err)
	}
	log.Printf("User: %d %s %s", result.User.ID, result.User.Username, result.User.Email)
}
