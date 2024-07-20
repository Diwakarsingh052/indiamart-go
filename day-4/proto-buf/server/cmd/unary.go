package main

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"server/gen/proto"
)

// User struct is a representation of user entity
type User struct {
	Name  string   `json:"name" validate:"required"`        // Name of the user, a required field
	Email string   `json:"email" validate:"required,email"` // Email ID of the user, should be in valid email format
	Roles []string `json:"roles" validate:"required"`       // Roles assigned to the user, a required field
}

func (userService) Signup(ctx context.Context, req *proto.SignupRequest) (*proto.SignupResponse, error) {
	nu := req.GetUser() // Fetching the user request sent by the client
	if nu == nil {
		return nil, status.Error(codes.Internal, "please provide required fields in correct format")
	}

	// Parsing the received request to our User struct
	u := User{
		Name:  nu.Name,
		Email: nu.Email,
		Roles: nu.Roles,
	}

	v := validator.New() // Creating a new validator instance

	// Validating the user struct according to the rules defined in the User struct.
	err := v.Struct(u)

	if err != nil {
		// If validation fails, return an error message with the error status.
		return nil, status.Error(codes.Internal, "please provide required fields in correct format")
	}

	fmt.Println(u)
	// Returning the Signup Response with the result showing the email ID with account created message.
	return &proto.SignupResponse{Result: u.Email + " account created"}, nil

}
