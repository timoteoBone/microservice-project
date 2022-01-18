package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"discordbot/main.go/go/pkg/mod/github.com/gorilla/mux@v1.8.0"

	"github.com/gorilla/mux"
	"github.com/timoteoBone/microservice-project/grpcService/pkg/entities"
	"github.com/timoteoBone/microservice-project/grpcService/pkg/errors"
	myerr "github.com/timoteoBone/microservice-project/grpcService/pkg/errors"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPSrv(endpoint Endpoints, logger log.Logger) http.Handler {
	rt := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse),
	}

	rt.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoint.CreateUs,
		decodeCreateUserReq,
		encodeCreateUserResp,
		options...,
	))

	rt.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoint.GetUs,
		decodeGetUserReq,
		encodeGetUserResp,
		options...,
	))

	rt.Methods("DELETE").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoint.DeleteUs,
		decodeDeleteRequest,
		encodeDeleteUserResponse,
		options,
	))
	return rt
}

func decodeCreateUserReq(ctx context.Context, r *http.Request) (interface{}, error) {

	var request entities.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, myerr.NewFieldsMissing()
	}

	return request, nil
}

func encodeCreateUserResp(ctx context.Context, wr http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(wr).Encode(response)
}

func decodeGetUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var request entities.GetUserRequest
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, myerr.NewFieldsMissing()
	}

	request.UserID = id
	return request, nil
}

func encodeGetUserResp(ctx context.Context, wr http.ResponseWriter, response interface{}) error {

	return json.NewEncoder(wr).Encode(response)
}

func decodeDeleteRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	//missing variable for the request
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		//return nil, myerr.NewFieldsMissing()
	}
	//assign id to request
	return nil, errors.New("missing implementation")
}

func encodeDeleteUserResponse(ctx context.Context, r http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(r).Encode(response)
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(errors.CustomToHttp(err))
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
	}
}
