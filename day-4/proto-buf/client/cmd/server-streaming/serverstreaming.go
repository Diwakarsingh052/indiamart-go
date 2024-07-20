package main

import (
	"client/gen/cmd/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
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

	req := &proto.GetPostsRequest{UserId: 101}
	stream, err := client.GetPosts(ctx, req)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		//receiving values from stream
		post, err := stream.Recv()

		//if the server has finished sending the request, we will quit
		if err == io.EOF {
			break
		}
		//any other kind of error would be caught here
		if err != nil {
			log.Println(err)
			return
		}
		select {
		case <-stream.Context().Done():
			fmt.Println("server cancelled")
		default:

		}
		fmt.Println("reading stream")
		//printing data received
		fmt.Println(post)
		fmt.Println()

	}

}
