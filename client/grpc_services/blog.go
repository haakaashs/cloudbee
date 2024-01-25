package grpcservices

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/haakaashs/cloudbee/client/models"
	"github.com/haakaashs/cloudbee/protos/blog"
)

func CreateBlogPost(client blog.BlogServiceClient, input models.Post) (int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		if cancel != nil {
			cancel()
		}
	}()

	res, err := client.CreateBlogPost(ctx, &blog.Post{
		Title:           input.Title,
		Content:         input.Content,
		Author:          input.Author,
		PublicationData: time.Now().String(),
		Tags:            input.Tags,
	})

	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	log.Printf("response from the server: %v\n", res.PostId)
	return res.PostId, nil
}

func ReadBlogPost(client blog.BlogServiceClient, id int32) (models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		if cancel != nil {
			cancel()
		}
	}()

	res, err := client.ReadBlogPost(ctx, &blog.Id{PostId: id})

	if err != nil {
		log.Println(err.Error())
		return models.Post{}, err
	}

	return models.Post{
		ID:              res.PostId,
		Title:           res.Title,
		Content:         res.Content,
		Author:          res.Author,
		PublicationData: res.PublicationData,
	}, nil
}

func UpdateBlogPost(client blog.BlogServiceClient, input models.Post) (models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		if cancel != nil {
			cancel()
		}
	}()

	if input.PublicationData != "" {
		cusErr := errors.New("error in input data publication should not be updated")
		log.Println(cusErr)
		return models.Post{}, cusErr
	}

	res, err := client.UpdateBlogPost(ctx, &blog.Post{
		Title:   input.Title,
		Content: input.Content,
		Author:  input.Author,
		Tags:    input.Tags,
	})

	if err != nil {
		log.Println(err.Error())
		return models.Post{}, err
	}

	return models.Post{
		ID:              res.PostId,
		Title:           res.Title,
		Content:         res.Content,
		Author:          res.Author,
		PublicationData: res.PublicationData,
	}, nil
}

func DeleteBlogPost(client blog.BlogServiceClient, id int32) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		if cancel != nil {
			cancel()
		}
	}()

	res, err := client.DeleteBlogPost(ctx, &blog.Id{PostId: id})

	if err != nil {
		log.Println(err.Error())
		return res.Message, err
	}

	return res.Message, nil
}
