package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"

	"grpc-test/user"
)

func main() {
	fmt.Println("server running")

	s := grpc.NewServer()

	user.RegisterUserServiceServer(s, &userServer{})

	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("could not listen to :50051 %v", err)
	}
	log.Fatal(s.Serve(l))
}

//NewUser(context.Context, *NewUserRequest) (*NewUserResponse, error)

type userServer struct{}

func (server *userServer) NewUser(ctx context.Context, req *user.NewUserRequest) (*user.NewUserResponse, error) {
	fmt.Println("Received request")
	fmt.Println(req.String())
	fmt.Println("a little time out")
	time.Sleep(time.Second * 2)
	fmt.Println("sending response")
	res := &user.NewUserResponse{
		Success: true,
	}
	return res, nil
}
