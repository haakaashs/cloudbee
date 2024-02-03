package grpcservices

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/haakaashs/cloudbee/client/models"
	"github.com/haakaashs/cloudbee/log"
	"github.com/haakaashs/cloudbee/protos/blog"
)

func CreateBlogPost(client blog.BlogServiceClient, input models.Post) (models.Post, error) {
	funcDesc := "CreateBlogPost"
	log.Generic.INFO("enter client " + funcDesc)

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
		PublicationDate: time.Now().String(),
		Tags:            input.Tags,
	})

	if err != nil {
		log.Generic.ERROR(err)
		return models.Post{}, err
	}

	input.ID = res.PostId
	log.Generic.INFO(fmt.Sprintf("response from the server: %v\n", res))
	log.Generic.INFO("exit client " + funcDesc)
	return input, nil
}

func ReadBlogPost(client blog.BlogServiceClient, id int32) (models.Post, error) {
	funcDesc := "ReadBlogPost"
	log.Generic.INFO("enter client " + funcDesc)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		if cancel != nil {
			cancel()
		}
	}()

	res, err := client.ReadBlogPost(ctx, &blog.Id{PostId: id})

	if err != nil {
		log.Generic.ERROR(err)
		return models.Post{}, err
	}
	log.Generic.INFO(fmt.Sprintf("response from the server: %v\n", res))
	log.Generic.INFO("exit client " + funcDesc)
	return models.Post{
		ID:              res.PostId,
		Title:           res.Title,
		Content:         res.Content,
		Author:          res.Author,
		PublicationDate: res.PublicationDate,
		Tags:            res.Tags,
	}, nil
}

func UpdateBlogPost(client blog.BlogServiceClient, input models.Post) (models.Post, error) {
	funcDesc := "UpdateBlogPost"
	log.Generic.INFO("enter client " + funcDesc)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		if cancel != nil {
			cancel()
		}
	}()

	if input.PublicationDate != "" {
		err := errors.New("error in input data publication should not be updated")
		log.Generic.ERROR(err)
		return models.Post{}, err
	}

	res, err := client.UpdateBlogPost(ctx, &blog.Post{
		PostId:  input.ID,
		Title:   input.Title,
		Content: input.Content,
		Author:  input.Author,
		Tags:    input.Tags,
	})

	if err != nil {
		log.Generic.ERROR(err)
		return models.Post{}, err
	}

	log.Generic.INFO(fmt.Sprintf("response from the server: %v\n", res))
	log.Generic.INFO("exit client " + funcDesc)
	return models.Post{
		ID:              res.PostId,
		Title:           res.Title,
		Content:         res.Content,
		Author:          res.Author,
		PublicationDate: res.PublicationDate,
		Tags:            res.Tags,
	}, nil
}

func DeleteBlogPost(client blog.BlogServiceClient, id int32) (string, error) {
	funcDesc := "DeleteBlogPost"
	log.Generic.INFO("enter client " + funcDesc)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		if cancel != nil {
			cancel()
		}
	}()

	res, err := client.DeleteBlogPost(ctx, &blog.Id{PostId: id})

	if err != nil {
		log.Generic.ERROR(err)
		return res.Message, err
	}

	log.Generic.INFO(fmt.Sprintf("response from the server: %v\n", res))
	log.Generic.INFO("exit client " + funcDesc)
	return res.Message, nil
}
