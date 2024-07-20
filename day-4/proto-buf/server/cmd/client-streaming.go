package main

import (
	"errors"
	"io"
	"log"
	"server/gen/proto"
)

// In client-streaming RPC, the client sends multiple messages/request to the server
// instead of a single request.
// The server sends back a single response to the client.

func (userService) CreatePost(stream proto.UserService_CreatePostServer) error {
	for {
		req, err := stream.Recv()
		//If the client has finished sending the request, we will quit
		if err == io.EOF {
			log.Println("stream ended")
			break
		}
		if err != nil {
			log.Println(err)
			return err
		}

		//if the request is cancelled then we would detect that using select
		//during the request if client close the connection we will know inside this select block

		select {
		case <-stream.Context().Done():
			log.Println("client cancelled the request")
			return errors.New("client disconnected")
		default:
			// Client is still connected
		}

		log.Printf("Received Create Post Requests: %v\n", req.GetPosts())
		//send the post for further processing, pass the context to the next call
		// if the ctx is cancelled , the other layer would also know and can rollback changes made
	}

	return stream.SendAndClose(&proto.CreatePostResponse{Result: "done"})
}
