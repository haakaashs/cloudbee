package grpcservices

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/haakaashs/cloudbee/client/models"
	"github.com/haakaashs/cloudbee/protos/blog"
)

var Client blog.BlogServiceClient

func CreatePost(w http.ResponseWriter, r *http.Request) {
	funcDesc := "CreatePost"
	log.Println("enter rest " + funcDesc)

	w.Header().Set("Content-Type", "application/json")
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	res, err := CreateBlogPost(Client, post)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	log.Println("exit rest " + funcDesc)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func ReadPost(w http.ResponseWriter, r *http.Request) {
	funcDesc := "ReadPost"
	log.Println("enter rest " + funcDesc)

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("errror in rest invalid id")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	res, err := ReadBlogPost(Client, int32(id))
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	log.Println("exit rest " + funcDesc)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	funcDesc := "UpdatePost"
	log.Println("enter rest " + funcDesc)

	w.Header().Set("Content-Type", "application/json")
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	res, err := UpdateBlogPost(Client, post)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	log.Println("exit rest " + funcDesc)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	funcDesc := "DeletePost"
	log.Println("enter rest " + funcDesc)

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("errror in rest invalid payload")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	res, err := DeleteBlogPost(Client, int32(id))
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	log.Println("exit rest " + funcDesc)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"` + res + `"}`))
}
