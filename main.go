package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	serviceId, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Error converting service ID: %v\n", err)
		panic(err)
	}
	if serviceId <= 0 {
		fmt.Printf("Invalid service ID: %d\n", serviceId)
		panic(fmt.Errorf("invalid service ID: %d", serviceId))
	}

	serviceName, err := GetServiceName(serviceId)
	if err != nil {
		fmt.Printf("Error getting service name: %v\n", err)

		panic(err)
	}

	fmt.Printf("Service name: %s\n is being deployed", serviceName)

	err = AddDeployment(serviceId)
	if err != nil {
		fmt.Printf("Error adding deployment: %v\n", err)

		panic(err)
	}

	fmt.Printf("Deployment added for service: %s\n", serviceName)
}
