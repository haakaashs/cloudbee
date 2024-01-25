package service

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/haakaashs/cloudbee/protos/blog"
)

type server struct {
	BlogPost map[int32]*blog.Post
	counter  int32
	mu       sync.Mutex
	blog.BlogServiceServer
}

func NewBlogServer() *server {
	return &server{}
}

func (s *server) CreateBlogPost(ctx context.Context, post *blog.Post) (*blog.Post, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	if out, ok := s.BlogPost[post.PostId]; ok {
		log.Println("error at server create blog post")
		return out, errors.New("post already exist")
	}

	s.counter++
	post.PostId = s.counter
	s.BlogPost[post.PostId] = post

	return post, nil
}

func (s *server) ReadBlogPost(ctx context.Context, id *blog.Id) (*blog.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var (
		post *blog.Post
		ok   bool
	)

	if post, ok = s.BlogPost[id.PostId]; !ok {
		log.Println("error at server read blog post")
		return nil, errors.New("post with id " + string(id.PostId) + " does not exist")
	}

	return post, nil
}

func (s *server) UpdateBlogPost(ctx context.Context, post *blog.Post) (*blog.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var (
		data *blog.Post
		ok   bool
	)

	if data, ok = s.BlogPost[post.PostId]; !ok {
		log.Println("error at server update blog post")
		return nil, errors.New("post with id " + string(post.PostId) + "does not exist")
	}

	post.PublicationData = data.PublicationData
	s.BlogPost[post.PostId] = post

	return post, nil
}

func (s *server) DeleteBlogPost(ctx context.Context, id *blog.Id) (*blog.DeleteResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.BlogPost[id.PostId]; !ok {
		log.Println("error at server delete blog post")
		return &blog.DeleteResponse{Message: "Failure"}, errors.New("post with id " + string(id.PostId) + " does not exist")
	}
	delete(s.BlogPost, id.PostId)
	return &blog.DeleteResponse{Message: "Success"}, nil
}
