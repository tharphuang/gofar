package main

import (
	"context"
	"fmt"

	pb "github.com/TharpHuang/gofar/generated/grpc"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("err:", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &pb.Request{Name: "QAQ"})
	if err != nil {
		fmt.Println(":", err)
		return
	}
	fmt.Println(r.Msg, r.Result)
}
