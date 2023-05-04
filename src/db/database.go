package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirestoreClient *firestore.Client

func InitFirebaseApp() {
	// Initialize Firebase app with service account credentials
	ctx := context.Background()

	credsFile := "C:/Users/Gabri/OneDrive/Documents/Projects/estudo-api/configPrivate/firebase.json"

	opt := option.WithCredentialsFile(credsFile)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing Firebase app: %v\n", err)
	}

	// Get Firestore client from Firebase app
	FirestoreClient, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error getting Firestore client: %v\n", err)
	}
}
