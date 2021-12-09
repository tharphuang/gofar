package main

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/structpb"

	pb "github.com/TharpHuang/gofar/generated/grpc"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (g *GreeterServer) SayHello(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	var any interface{} = "123"

	res, _ := structpb.NewValue(any)
	return &pb.Response{
		Msg:    "he " + in.Name,
		Result: res,
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", "localhost:50051")
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &GreeterServer{})
	reflection.Register(s)
	s.Serve(lis)
}
