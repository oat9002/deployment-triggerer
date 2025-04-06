package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var firebaseAppOnce sync.Once
var fireStoreOnce sync.Once
var firebaseApp *firebase.App
var firestoreClient *firestore.Client

func initializeFirebase() (*firebase.App, error) {
	var err error

	if firebaseApp != nil {
		return firebaseApp, nil
	}

	firebaseAppOnce.Do(func() {
		serviceAccountKeyJSON := os.Getenv("DEPLOYMENT_TRIGGERER_FIREBASE_SERVICE_ACCOUNT_KEY_JSON")
		if serviceAccountKeyJSON == "" {
			err = fmt.Errorf("DEPLOYMENT_TRIGGERER_FIREBASE_SERVICE_ACCOUNT_KEY_JSON is required")
		}

		credential := option.WithCredentialsJSON([]byte(serviceAccountKeyJSON))
		firebaseApp, err = firebase.NewApp(context.Background(), nil, credential)
	})

	return firebaseApp, err
}

func GetFireStoreClient() (*firestore.Client, error) {
	firebaseApp, err := initializeFirebase()
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	if firestoreClient != nil {
		return firestoreClient, nil
	}

	fireStoreOnce.Do(func() {
		firestoreClient, err = firebaseApp.Firestore(context.Background())
	})

	return firestoreClient, err
}

func GetDocumentById(client *firestore.Client, collection string, id string) (map[string]interface{}, error) {
	doc, err := client.Collection(collection).Doc(id).Get(context.Background())

	if err != nil {
		return nil, err
	}

	return doc.Data(), nil
}

func AddDocument(client *firestore.Client, collection string, data interface{}) error {
	_, _, err := client.Collection(collection).Add(context.Background(), data)

	if err != nil {
		return err
	}

	return nil
}
