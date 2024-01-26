package main

import (
	"net"

	"github.com/haakaashs/cloudbee/log"
	"github.com/haakaashs/cloudbee/protos/blog"
	"github.com/haakaashs/cloudbee/server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	defer log.File.Close()

	const port = ":50001"

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Generic.ERROR(err)
		log.Generic.FATAL("unable to listen on port " + port)
	}

	ser := grpc.NewServer()

	blog.RegisterBlogServiceServer(ser, service.NewBlogServer())
	reflection.Register(ser)

	log.Generic.INFO("server listening on port " + port)
	if err := ser.Serve(listener); err != nil {
		log.Generic.ERROR(err)
		log.Generic.FATAL("unable to start the server " + port)
	}
}
