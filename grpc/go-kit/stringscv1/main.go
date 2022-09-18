package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

//1.定义服务接口和实现
// StringService provides operations on strings.
type IStringService interface {
	UpperCase(string) (string, error)
	Count(string) int
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

type StringService struct{}

func (ss *StringService) UpperCase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (ss *StringService) Count(s string) int {
	return len(s)
}

//2.endpoint 方法相关request response
type UpperCaseRequest struct {
	S string `json:"s"`
}

type UpperCaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

type CountRequest struct {
	S string `json:"s"`
}

type CountResponse struct {
	V int `json:"v"`
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func GenUpperCaseEndpoint(svc IStringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpperCaseRequest)
		v, err := svc.UpperCase(req.S)
		if err != nil {
			return UpperCaseResponse{v, err.Error()}, nil
		}
		return UpperCaseResponse{v, ""}, nil
	}
}

func GenCountEndpoint(svc IStringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CountRequest)
		v := svc.Count(req.S)
		return CountResponse{v}, nil
	}
}

//transport 数据格式处理编解码
func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UpperCaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	svc := &StringService{}

	uppercaseHandler := httptransport.NewServer(
		GenUpperCaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		GenCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
