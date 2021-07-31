package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/gorilla/mux"
)

func healthDecodeRequest(c context.Context, request *http.Request) (interface{}, error) {
	return struct{}{}, nil
}

func healthEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}

func factorialDecodeRequest(c context.Context, request *http.Request) (interface{}, error) {
	g := factorialRequest{}
	u := strings.Split(request.URL.Path, "/")
	if len(u) > 2 {
		Num, err := strconv.Atoi(u[2])
		if err != nil {
			log.Println(err)
			g.Num = 0
			return g, nil
		}
		g.Num = Num
	}
	return g, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	// All of our response objects are JSON serializable, so we just do that.
	w.Header().Set("Content-Type", "application/hal+json")
	return json.NewEncoder(w).Encode(response)
}

// MakeHttpHandler make http handler use mux
func MakeHttpHandler(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()

	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(kithttp.DefaultErrorEncoder),
	}
	r.Methods("GET").PathPrefix("/health").Handler(kithttp.NewServer(
		endpoints.HealthEndpoint,
		healthDecodeRequest,
		healthEncodeResponse,
		options...,
	))
	r.Methods("GET").PathPrefix("/factorial").Handler(kithttp.NewServer(
		endpoints.FactorialEndpoint,
		factorialDecodeRequest,
		encodeResponse,
		options...,
	))
	return r
}
