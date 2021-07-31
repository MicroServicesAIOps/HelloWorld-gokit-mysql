package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	HealthEndpoint    endpoint.Endpoint
	FactorialEndpoint endpoint.Endpoint
}

type healthRequest struct {
	//
}

type healthResponse struct {
	Health []Health `json:"health"`
}

type factorialRequest struct {
	Num int `json:"num"`
}

type factorialResponse struct {
	Factorial Factorial `json:"facval"`
}

func MakeFactorialEndpoint(s IMyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(factorialRequest)
		factorial := s.Factorial(int(req.Num))
		return factorialResponse{Factorial: factorial}, nil
	}
}

func MakeHealthEndpoint(s IMyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		health := s.Health()
		return healthResponse{Health: health}, nil
	}
}

func MakeEndpoints(s IMyService) Endpoints {
	health := MakeHealthEndpoint(s)
	factorial := MakeFactorialEndpoint(s)
	endpoints := Endpoints{
		HealthEndpoint:    health,
		FactorialEndpoint: factorial,
	}
	return endpoints
}
