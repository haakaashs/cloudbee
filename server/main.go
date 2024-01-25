package main

import (
	"log"
	"net"

	"github.com/haakaashs/cloudbee/protos/blog"
	"github.com/haakaashs/cloudbee/server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	const port = ":50001"

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("unable to listen on port: %s\n", err.Error())
	}

	ser := grpc.NewServer()

	blog.RegisterBlogServiceServer(ser, service.NewBlogServer())
	reflection.Register(ser)

	log.Printf("server listening on port: %s\n", port)
	if err := ser.Serve(listener); err != nil {
		log.Printf("unable to start the server: %s\n", port)
	}

	defer func() {
		if recover := recover(); recover != nil {
			log.Println("recovered form panic")
		}
	}()
}
