package repository

import (
	"context"
	"encoding/json"
	"github.com/gabriel/estudo-api/src/db"
	"github.com/gabriel/estudo-api/src/models"
	"google.golang.org/api/iterator"
	"net/http"
)

func GetAllClients(ctx context.Context) ([]models.Clientstruct, error) {
	iter := db.FirestoreClient.Collection("clients").Documents(ctx)

	var clients []models.Clientstruct

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var client models.Clientstruct
		err = doc.DataTo(&client)
		if err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	var client []models.Clientstruct
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request payload"})
		return
	}

	_, _, err = db.FirestoreClient.Collection("clients").Add(context.Background(), client)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "could not create client"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "client created successfully"})
	return
}
