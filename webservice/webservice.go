package webservice

import (
	"context"
	"fmt"
	"krispogram-grpc/infrastructure/db"
	"krispogram-grpc/pb"
	"math/rand"
)

type Server struct {
	dbInteractor *db.DbInteractor
}

func NewServer(dbInteractor *db.DbInteractor) *Server {
	server := Server{dbInteractor: dbInteractor}
	server.dbInteractor.Connect()
	return &server
}

func (s *Server) Create(ctx context.Context, post *pb.NewPost) (*pb.Post, error) {
	fmt.Printf("Receive message post from client with id = %d, topic - %s", post.AuthorId, post.Topic)
	savedPost := pb.Post{
		Id:       uint64(rand.Int()),
		AuthorId: post.AuthorId,
		Topic:    post.Topic,
		Body:     post.Body,
	}
	err := s.dbInteractor.Create(&savedPost)
	if err != nil {
		return nil, err
	}
	return &savedPost, nil
}

func (s *Server) GetById(ctx context.Context, req *pb.GetPostByIdReq) (*pb.Post, error) {
	fmt.Printf("Receive message request post with id = %d", req.Id)
	res, err := s.dbInteractor.GetById(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Server) GetByAuthorId(ctx context.Context, req *pb.GetPostsByAuthorIdReq) (*pb.GetPostsByAuthorIdResp, error) {
	fmt.Printf("Receive message request posts from author with id = %d", req.AuthorId)
	res, err := s.dbInteractor.GetByAuthorId(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
