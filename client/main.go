package main

import (
	"context"
	"fmt"
	"os"

	"grpc-test/user"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())

	if err != nil {
		fmt.Printf("could not connect to backend: %v\n", err)
		os.Exit(1)
	}

	client := user.NewUserServiceClient(conn)
	addUser(client)

}

func addUser(client user.UserServiceClient) {
	newUser := &user.User{
		Name:  "Michael",
		Email: "m@gmail.com",
	}
	req := user.NewUserRequest{
		User: newUser,
	}

	res, err := client.NewUser(context.Background(), &req)
	if err != nil {
		fmt.Printf("error while calling GreetManyTimes RPC: %v", err)
	}
	fmt.Println(res.String())
}
