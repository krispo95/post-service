package main

import (
	"fmt"
	"google.golang.org/grpc"
	"krispogram-grpc/pb"
	"krispogram-grpc/webservice"
	"log"
	"net"
	"runtime/debug"
)

func main() {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Printf(fmt.Sprintf("Stack trace\n%s", string(debug.Stack())), "level", "panic")
		}
	}()
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := webservice.Server{}

	grpcServer := grpc.NewServer()
	pb.RegisterPostServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
