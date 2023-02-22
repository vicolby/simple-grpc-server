package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/vicolby/simple-grpc-server/gen/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTestApiClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	r, err := c.Echo(ctx, &pb.EchoRequest{Message: "Hello World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r)

}
