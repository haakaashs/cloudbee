package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	grpcservices "github.com/haakaashs/cloudbee/client/grpc_services"
	"github.com/haakaashs/cloudbee/protos/blog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	url  = ":50001"
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connection to server failed %v\n", err)
	}

	defer conn.Close()

	log.Fatalf("client listening at port %s\n", url)

	grpcservices.Client = blog.NewBlogServiceClient(conn)
	// we can directly hit the server using client

	// or we can indirectly hit the server using client using rest API
	server := mux.NewRouter()

	server.HandleFunc("/blog", grpcservices.CreatePost).Methods("POST")
	server.HandleFunc("/blog/{id}", grpcservices.ReadPost).Methods("GET")
	server.HandleFunc("/blog", grpcservices.UpdatePost).Methods("PUT")
	server.HandleFunc("/blog/{id}", grpcservices.DeletePost).Methods("DELETE")

	log.Println("rest server started listening at port :8080......")
	if err = http.ListenAndServe(port, server); err != nil {
		log.Fatalf("error starting rest web server at port %s\n", port)
	}

	defer func() {
		if recover := recover(); recover != nil {
			log.Println("recovered form panic")
		}
	}()
}
