package webservice

import (
	"context"
	"fmt"
	"krispogram-grpc/pb"
	"math/rand"
)

type Server struct{}

func (s *Server) Create(ctx context.Context, post *pb.NewPost) (*pb.Post, error) {
	fmt.Printf("Receive message post from client with id = %d, topic - %s", post.AuthorId, post.Topic)
	savedPost := pb.Post{
		Id:       uint64(rand.Int()),
		AuthorId: post.AuthorId,
		Topic:    post.Topic,
		Body:     post.Body,
	}
	return &savedPost, nil
}

func (s *Server) GetById(ctx context.Context, req *pb.GetPostByIdReq) (*pb.Post, error) {
	fmt.Printf("Receive message request post with id = %d", req.Id)
	return &pb.Post{}, nil
}

func (s *Server) GetByAuthorId(ctx context.Context, req *pb.GetPostsByAuthorIdReq) (*pb.GetPostsByAuthorIdResp, error) {
	fmt.Printf("Receive message request posts from author with id = %d", req.AuthorId)
	return &pb.GetPostsByAuthorIdResp{}, nil
}
