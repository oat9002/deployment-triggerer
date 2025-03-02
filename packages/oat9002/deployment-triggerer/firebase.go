package main

import (
	"context"
	"log"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

var once sync.Once
var firebaseApp *firebase.App

func initializeFirebase() (*firebase.App, error) {
	var err error

	if firebaseApp != nil {
		return firebaseApp, nil
	}

	once.Do(func() {
		firebaseApp, err = firebase.NewApp(context.Background(), nil)
	})

	return firebaseApp, err
}

func GetFireStoreClient() (*firestore.Client, error) {
	firebaseApp, err := initializeFirebase()
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return firebaseApp.Firestore(context.Background())
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
