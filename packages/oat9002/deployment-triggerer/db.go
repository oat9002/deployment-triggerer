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
	ServiceId int32     `json:"service_id"`
	CreatedAt time.Time `json:"created_at"`
}

func AddDeployment(serviceId int32) error {
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
	}

	err = AddDocument(client, deploymentsCollection, deploymentData)

	return err
}
