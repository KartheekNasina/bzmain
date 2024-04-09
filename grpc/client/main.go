package main

import (
	"context"
	"log"
	"time"

	pb "github.com/vivekbnwork/bz-backend/bz-main/grpc/proto/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	startTime := time.Now() // Capture the start time

	r, err := c.GetAllUsers(ctx, &pb.GetAllUsersRequest{})
	if err != nil {
		log.Fatalf("could not retrieve users: %v", err)
	}

	endTime := time.Now()                  // Capture the end time
	responseTime := endTime.Sub(startTime) // Calculate the response time

	// Print the received users.
	for _, user := range r.Users {
		log.Printf("User ID: %s, Name: %s", user.Id, user.Name)
	}

	// Print the response time.
	log.Printf("Response Time: %v", responseTime)
}
