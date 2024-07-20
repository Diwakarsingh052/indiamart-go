package main

import (
	"client/gen/cmd/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	dialOpts := []grpc.DialOption{
		// WithTransportCredentials specifies the transport credentials for the connection
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient("localhost:5001", dialOpts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	// Ensure to close the connection when the main function finishes
	defer conn.Close()

	// Instantiate a client for the UserService service
	client := proto.NewUserServiceClient(conn)

	// Create a context with a 10-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// Cancel the context when the main function finishes
	defer cancel()

	// Create a SignupRequest object
	req := &proto.SignupRequest{
		User: &proto.User{
			// Set the user name to John
			Name: "John",
			// Set the user email
			Email: "johnemail.com",
			// Set the password for the user
			Password: "abc",
			// Set the user roles to ADMIN and USER
			Roles: []string{"ADMIN", "USER"},
		},
	}
	resp, err := client.Signup(ctx, req)
	if err != nil {
		log.Println(err)
		return
	}

	// Log the result of the Signup request
	log.Println(resp.Result)

}
