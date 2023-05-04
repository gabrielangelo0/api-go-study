package controllers

import (
	"context"
	"encoding/json"
	"github.com/gabriel/estudo-api/src/repository"
	"net/http"
)

func CreateClient(w http.ResponseWriter, r *http.Request) {}

func ListClients(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	clients, err := repository.GetAllClients(ctx)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "could not get clients"})
		return
	}

	responseJSON, err := json.Marshal(clients)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "could not serialize response"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func ShowClient(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Showing a Client"))
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating a Client"))
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting a Client"))
}
