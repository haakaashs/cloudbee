package test

import (
	"context"
	"log"
	"testing"

	"github.com/go-playground/assert"
	"github.com/haakaashs/cloudbee/protos/blog"
	"github.com/haakaashs/cloudbee/server/service"
)

var title = "new title"

func TestCreateBlogPost(t *testing.T) {
	tests := []struct {
		Desc     string
		Input    *blog.Post
		Expected interface{}
		Ctx      context.Context
	}{
		{
			Desc:     "success test case",
			Input:    &blog.Post{Title: title},
			Expected: &blog.Post{Title: title},
			Ctx:      context.Background(),
		}, {
			Desc:     "failure test case",
			Input:    &blog.Post{PostId: 1},
			Expected: "post already exist",
			Ctx:      context.Background(),
		},
	}
	serverObj := service.NewBlogServer()
	for _, tt := range tests {
		log.Println(tt.Desc)
		t.Run(tt.Desc, func(t *testing.T) {
			res, err := serverObj.CreateBlogPost(tt.Ctx, tt.Input)

			if err != nil {
				assert.Equal(t, err.Error(), tt.Expected.(string))
				return
			}
			assert.Equal(t, res.Title, tt.Expected.(*blog.Post).Title)

		})
	}
}

func TestReadBlogPost(t *testing.T) {
	tests := []struct {
		Desc     string
		Input    *blog.Id
		Expected interface{}
		Ctx      context.Context
	}{
		{
			Desc:     "success test case",
			Input:    &blog.Id{PostId: 1},
			Expected: &blog.Post{PostId: 1},
			Ctx:      context.Background(),
		}, {
			Desc:     "failure test case",
			Input:    &blog.Id{PostId: 2},
			Expected: "post id does not exist",
			Ctx:      context.Background(),
		},
	}
	serverObj := service.NewBlogServer()
	serverObj.CreateBlogPost(context.Background(), &blog.Post{})
	for _, tt := range tests {
		log.Println(tt.Desc)
		t.Run(tt.Desc, func(t *testing.T) {
			res, err := serverObj.ReadBlogPost(tt.Ctx, tt.Input)
			if err != nil {
				log.Println(err.Error())
				assert.Equal(t, err.Error(), tt.Expected.(string))
				return
			}
			assert.Equal(t, res.PostId, tt.Expected.(*blog.Post).PostId)
		})
	}
}

func TestUpdateBlogPost(t *testing.T) {
	tests := []struct {
		Desc     string
		Input    *blog.Post
		Expected interface{}
		Ctx      context.Context
	}{
		{
			Desc:     "success test case",
			Input:    &blog.Post{PostId: 1},
			Expected: &blog.Post{PostId: 1},
			Ctx:      context.Background(),
		}, {
			Desc:     "failure test case",
			Input:    &blog.Post{PostId: 2},
			Expected: "post id does not exist",
			Ctx:      context.Background(),
		},
	}
	serverObj := service.NewBlogServer()
	serverObj.CreateBlogPost(context.Background(), &blog.Post{})
	for _, tt := range tests {
		log.Println(tt.Desc)
		t.Run(tt.Desc, func(t *testing.T) {
			res, err := serverObj.UpdateBlogPost(tt.Ctx, tt.Input)
			if err != nil {
				assert.Equal(t, err.Error(), tt.Expected.(string))
				return
			}
			assert.Equal(t, res.PostId, tt.Expected.(*blog.Post).PostId)
		})
	}
}

func TestDeleteBlogPost(t *testing.T) {
	tests := []struct {
		Desc     string
		Input    *blog.Id
		Expected *blog.DeleteResponse
		Ctx      context.Context
	}{
		{
			Desc:     "success test case",
			Input:    &blog.Id{PostId: 1},
			Expected: &blog.DeleteResponse{Message: "Success"},
			Ctx:      context.Background(),
		}, {
			Desc:     "failure test case",
			Input:    &blog.Id{PostId: 2},
			Expected: &blog.DeleteResponse{Message: "Failure"},
			Ctx:      context.Background(),
		},
	}
	serverObj := service.NewBlogServer()
	serverObj.CreateBlogPost(context.Background(), &blog.Post{})
	for _, tt := range tests {
		log.Println(tt.Desc)
		t.Run(tt.Desc, func(t *testing.T) {
			res, err := serverObj.DeleteBlogPost(tt.Ctx, tt.Input)
			if err != nil {
				assert.Equal(t, res.Message, tt.Expected.Message)
				assert.Equal(t, err.Error(), "post id does not exist")
				return
			}
			assert.Equal(t, res.Message, tt.Expected.Message)
		})
	}
}
