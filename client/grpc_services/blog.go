package grpcservices

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/haakaashs/cloudbee/client/models"
	"github.com/haakaashs/cloudbee/protos/blog"
)

func CreateBlogPost(client blog.BlogServiceClient, input models.Post) (models.Post, error) {
	funcDesc := "CreateBlogPost"
	log.Println("enter rest " + funcDesc)

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
		return models.Post{}, err
	}

	input.ID = res.PostId
	log.Printf("response from the server: %v\n", res)
	log.Println("exit rest " + funcDesc)
	return input, nil
}

func ReadBlogPost(client blog.BlogServiceClient, id int32) (models.Post, error) {
	funcDesc := "ReadBlogPost"
	log.Println("enter rest " + funcDesc)

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
	log.Printf("response from the server: %v\n", res)
	log.Println("exit rest " + funcDesc)
	return models.Post{
		ID:              res.PostId,
		Title:           res.Title,
		Content:         res.Content,
		Author:          res.Author,
		PublicationData: res.PublicationData,
		Tags:            res.Tags,
	}, nil
}

func UpdateBlogPost(client blog.BlogServiceClient, input models.Post) (models.Post, error) {
	funcDesc := "UpdateBlogPost"
	log.Println("enter rest " + funcDesc)

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
		PostId:  input.ID,
		Title:   input.Title,
		Content: input.Content,
		Author:  input.Author,
		Tags:    input.Tags,
	})

	if err != nil {
		log.Println(err.Error())
		return models.Post{}, err
	}

	log.Printf("response from the server: %v\n", res)
	log.Println("exit rest " + funcDesc)
	return models.Post{
		ID:              res.PostId,
		Title:           res.Title,
		Content:         res.Content,
		Author:          res.Author,
		PublicationData: res.PublicationData,
		Tags:            res.Tags,
	}, nil
}

func DeleteBlogPost(client blog.BlogServiceClient, id int32) (string, error) {
	funcDesc := "DeleteBlogPost"
	log.Println("enter rest " + funcDesc)

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

	log.Printf("response from the server: %v\n", res)
	log.Println("exit rest " + funcDesc)
	return res.Message, nil
}
