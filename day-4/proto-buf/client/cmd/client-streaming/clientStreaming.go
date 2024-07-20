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

	// Simulate first batch of posts
	batch1 := []*proto.Post{
		{
			Title:  "The Science of Design",
			Author: "Author 1",
			Body:   "Body of post 1",
		},
		{
			Title:  "The Politics of Power",
			Author: "Author 2",
			Body:   "Body of post 2",
		},
		{
			Title:  "The Art of Programming",
			Author: "Author 3",
			Body:   "Body of post 3",
		},
	}
	req := &proto.CreatePostRequest{Posts: batch1}

	// Call the CreatePost method of the client and get a reply stream, along with any error
	stream, err := client.CreatePost(ctx)
	// Attempt to send CreatePost request through the stream
	err = stream.Send(req)

	// If there was an error in sending  request, then log and exit
	if err != nil {
		log.Fatalf("Failed to createPost request: %v", err)
	}
	time.Sleep(5 * time.Second)
	// Simulate second batch of posts
	batch2 := []*proto.Post{
		{
			Title:  "Post 11",
			Author: "Author 1",
			Body:   "Body of post 1",
		},
		{
			Title:  "Post 21",
			Author: "Author 2",
			Body:   "Body of post 2",
		},
		{
			Title:  "Post 31",
			Author: "Author 3",
			Body:   "Body of post 3",
		},
	}

	// Put the second batch in a CreatePostRequest object
	req = &proto.CreatePostRequest{Posts: batch2}

	// Attempt to send CreatePost request through the stream
	err = stream.Send(req)

	// If there was an error in sending  request, then log and exit
	if err != nil {
		log.Fatalf("Failed to createPost request: %v", err)
	}

	// Close the client streaming and receive the server's response
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive response: %v", err)

	}
	// Log the response from the server
	log.Printf("Response: %s", resp.Result)
}
