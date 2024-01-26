package grpcservices

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/haakaashs/cloudbee/client/models"
	"github.com/haakaashs/cloudbee/log"
	"github.com/haakaashs/cloudbee/protos/blog"
)

var Client blog.BlogServiceClient

func CreatePost(w http.ResponseWriter, r *http.Request) {
	funcDesc := "CreatePost"
	log.Generic.INFO("enter rest " + funcDesc)

	w.Header().Set("Content-Type", "application/json")
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Generic.ERROR(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	res, err := CreateBlogPost(Client, post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	log.Generic.INFO("exit rest " + funcDesc)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func ReadPost(w http.ResponseWriter, r *http.Request) {
	funcDesc := "ReadPost"
	log.Generic.INFO("enter rest " + funcDesc)

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Generic.ERROR(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	res, err := ReadBlogPost(Client, int32(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	log.Generic.INFO("exit rest " + funcDesc)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	funcDesc := "UpdatePost"
	log.Generic.INFO("enter rest " + funcDesc)

	w.Header().Set("Content-Type", "application/json")
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Generic.ERROR(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	res, err := UpdateBlogPost(Client, post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	log.Generic.INFO("exit rest " + funcDesc)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	funcDesc := "DeletePost"
	log.Generic.INFO("enter rest " + funcDesc)

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Generic.ERROR(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}
	res, err := DeleteBlogPost(Client, int32(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	log.Generic.INFO("exit rest " + funcDesc)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"` + res + `"}`))
}
