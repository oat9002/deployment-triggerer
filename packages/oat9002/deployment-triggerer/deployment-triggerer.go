package main

import (
	"fmt"
)

type Request struct {
	ServiceId int32 `json:"serviceId"`
}

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func Main(in Request) (*Response, error) {
	if in.ServiceId <= 0 {
		return nil, fmt.Errorf("serviceId must be greater than 0")
	}

	err := AddDeployment(in.ServiceId)
	if err != nil {
		return nil, err
	}

	serviceName, err := GetServiceName(in.ServiceId)
	if err != nil {
		return nil, err
	}

	return &Response{
		Body: fmt.Sprintf("Service %s is added for deployment sucessfully!", serviceName),
	}, nil
}
