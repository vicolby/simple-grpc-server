package main

import (
	"context"
	"net"

	pb "github.com/vicolby/simple-grpc-server/gen/proto"
	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoRequest, error) {
	return req, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterTestApiServer(s, &testApiServer{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
