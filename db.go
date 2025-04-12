package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

const deploymentsCollection = "deployments"
const servicesCollection = "services"

type Deployment struct {
	ServiceId int       `firestore:"service_id"`
	CreatedAt time.Time `firestore:"created_at"`
	IsSuccess bool      `firestore:"is_success"`
}

func GetServiceName(serviceId int) (string, error) {
	client, err := GetFireStoreClient()
	if err != nil {
		log.Fatalf("error getting firestore client: %v\n", err)
		return "", err
	}

	data, err := GetDocumentById(client, servicesCollection, strconv.Itoa(int(serviceId)))
	if err != nil {
		log.Fatalf("error getting service: %v\n", err)
		return "", err
	}

	return data["name"].(string), nil
}

func AddDeployment(serviceId int) error {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Fatalf("error loading location: %v\n", err)
		return err
	}

	client, err := GetFireStoreClient()
	if err != nil {
		log.Fatalf("error getting firestore client: %v\n", err)
		return err
	}

	data, err := GetDocumentById(client, servicesCollection, strconv.Itoa(int(serviceId)))
	if err != nil {
		log.Fatalf("error getting service: %v\n", err)
		return err
	}

	if data == nil {
		log.Fatalf("service not found: %v\n", serviceId)

		return fmt.Errorf("service not found: %v", serviceId)
	}

	deploymentData := Deployment{
		ServiceId: serviceId,
		CreatedAt: time.Now().In(loc),
		IsSuccess: false,
	}

	err = AddDocument(client, deploymentsCollection, deploymentData)

	return err
}
