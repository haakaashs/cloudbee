package service

import (
	"context"
	"errors"
	"sync"

	"github.com/haakaashs/cloudbee/log"
	"github.com/haakaashs/cloudbee/protos/blog"
)

type server struct {
	BlogPost map[int32]*blog.Post
	counter  int32
	mu       sync.Mutex
	blog.BlogServiceServer
}

func NewBlogServer() *server {
	return &server{
		BlogPost: make(map[int32]*blog.Post),
	}
}

func (s *server) CreateBlogPost(ctx context.Context, post *blog.Post) (*blog.Post, error) {
	funcDesc := "CreateBlogPost"
	log.Generic.INFO("enter rest " + funcDesc)

	s.mu.Lock()
	defer s.mu.Unlock()

	if out, ok := s.BlogPost[post.PostId]; ok {
		err := errors.New("post already exist")
		log.Generic.ERROR(err)
		return out, err
	}

	s.counter++
	post.PostId = s.counter
	s.BlogPost[post.PostId] = post

	log.Generic.INFO("exit rest " + funcDesc)
	return post, nil
}

func (s *server) ReadBlogPost(ctx context.Context, id *blog.Id) (*blog.Post, error) {
	funcDesc := "ReadBlogPost"
	log.Generic.INFO("enter rest " + funcDesc)

	s.mu.Lock()
	defer s.mu.Unlock()

	var (
		post *blog.Post
		ok   bool
	)

	if post, ok = s.BlogPost[id.PostId]; !ok {
		err := errors.New("post id does not exist")
		log.Generic.ERROR(err)
		return nil, err
	}

	log.Generic.INFO("exit rest " + funcDesc)
	return post, nil
}

func (s *server) UpdateBlogPost(ctx context.Context, post *blog.Post) (*blog.Post, error) {
	funcDesc := "UpdateBlogPost"
	log.Generic.INFO("enter rest " + funcDesc)

	s.mu.Lock()
	defer s.mu.Unlock()

	var (
		data *blog.Post
		ok   bool
	)

	if data, ok = s.BlogPost[post.PostId]; !ok {
		err := errors.New("post id does not exist")
		log.Generic.ERROR(err)
		return nil, err
	}

	post.PublicationData = data.PublicationData
	s.BlogPost[post.PostId] = post

	log.Generic.INFO("exit rest " + funcDesc)
	return post, nil
}

func (s *server) DeleteBlogPost(ctx context.Context, id *blog.Id) (*blog.DeleteResponse, error) {
	funcDesc := "DeleteBlogPost"
	log.Generic.INFO("enter rest " + funcDesc)

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.BlogPost[id.PostId]; !ok {
		err := errors.New("post id does not exist")
		log.Generic.ERROR(err)
		return &blog.DeleteResponse{Message: "Failure"}, err
	}
	delete(s.BlogPost, id.PostId)
	log.Generic.INFO("exit rest " + funcDesc)
	return &blog.DeleteResponse{Message: "Success"}, nil
}
