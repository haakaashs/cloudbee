package main

import (
	"net/http"

	"github.com/gorilla/mux"
	grpcservices "github.com/haakaashs/cloudbee/client/grpc_services"
	"github.com/haakaashs/cloudbee/log"
	"github.com/haakaashs/cloudbee/protos/blog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcPort = ":50001"
	restPort = ":8080"
)

func main() {
	conn, err := grpc.Dial(grpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Generic.ERROR(err)
		log.Generic.FATAL("connection to server failed")
	}

	defer conn.Close()

	log.Generic.INFO("client listening at port " + grpcPort)

	grpcservices.Client = blog.NewBlogServiceClient(conn)
	// we can directly hit the server using client

	// or we can indirectly hit the server using client using rest API
	server := mux.NewRouter()

	server.HandleFunc("/blog", grpcservices.CreatePost).Methods("POST")
	server.HandleFunc("/blog/{id}", grpcservices.ReadPost).Methods("GET")
	server.HandleFunc("/blog", grpcservices.UpdatePost).Methods("PUT")
	server.HandleFunc("/blog/{id}", grpcservices.DeletePost).Methods("DELETE")

	log.Generic.INFO("rest server started listening at port " + restPort)
	if err = http.ListenAndServe(restPort, server); err != nil {
		log.Generic.ERROR(err)
		log.Generic.FATAL("error starting rest web server at port " + restPort)
	}

	defer func() {
		if recover := recover(); recover != nil {
			log.Generic.WARN(recover)
			log.Generic.WARN("recovered form panic")
		}
	}()
}
