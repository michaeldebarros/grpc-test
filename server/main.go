package main

import (
	"fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"

	"/Users/michaeldebarros/Projects/grpc-test/user"
)

type server struct{}

func main() {
	fmt.Println("server running")

	s := grpc.NewServer()

	user.RegisterNewUserServiceServer(s, &server)

	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("could not listen to :8888: %v", err)
	}
	log.Fatal(s.Serve(l))
}
