package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/timoteoBone/microservice-project/grpcService/pkg/entities"
	myerr "github.com/timoteoBone/microservice-project/grpcService/pkg/errors"
	util "github.com/timoteoBone/microservice-project/httpService/pkg/utils"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPSrv(endpoint Endpoints, logger log.Logger) http.Handler {
	rt := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse),
	}

	rt.Methods("POST").Path(util.CreateUserPath).Handler(httptransport.NewServer(
		endpoint.CreateUs,
		decodeCreateUserReq,
		encodeCreateUserResp,
		options...,
	))

	rt.Methods("GET").Path(util.GetUserPath).Handler(httptransport.NewServer(
		endpoint.GetUs,
		decodeGetUserReq,
		encodeGetUserResp,
		options...,
	))

	rt.Methods("DELETE").Path(util.DeleteUserPath).Handler(httptransport.NewServer(
		endpoint.DeleteUs,
		decodeDeleteRequest,
		encodeDeleteUserResponse,
		options...,
	))

	rt.Methods("PUT").Path(util.PutUserPath).Handler(httptransport.NewServer(
		endpoint.UpdateUs,
		decodeUpdateUserRequest,
		encodeUpdateUserResponse,
		options...,
	))
	return rt
}

func decodeCreateUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var request entities.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, myerr.NewBadRequest()
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
		return nil, myerr.NewBadRequest()
	}

	request.UserID = id
	return request, nil
}

func encodeGetUserResp(ctx context.Context, wr http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(wr).Encode(response)
}

func decodeDeleteRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request entities.DeleteUserRequest
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, myerr.NewBadRequest()
	}
	request.UserId = id
	return request, nil
}

func encodeDeleteUserResponse(ctx context.Context, r http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(r).Encode(response)
}

func decodeUpdateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {

	request := entities.UpdateUserRequest{}
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, myerr.NewBadRequest()
	}
	request.Id = id

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, myerr.NewBadRequest()
	}

	return request, nil

}

func encodeUpdateUserResponse(ctx context.Context, rw http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(rw).Encode(response)
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(myerr.CustomToHttp(err))
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
	}
}
