package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"krispogram-grpc/pb"
	"time"
)

func main() {
	serverAddr := "localhost:9000"
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("error while starting connection, err - %v", err)
	}
	defer conn.Close()
	client := pb.NewPostServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()
	post, err := client.Create(ctx, &pb.NewPost{
		AuthorId: 1,
		Topic:    "topic1",
		Body:     "body1",
	})
	if err != nil {
		panic(fmt.Errorf("can't create post: %w", err))
	}
	fmt.Println(post.String())
}
