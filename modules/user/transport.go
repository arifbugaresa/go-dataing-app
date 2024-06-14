package user

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(endpoints Endpoints, r *mux.Router) http.Handler {
	r.Handle("/login", httptransport.NewServer(
		endpoints.LoginEndpoint,
		decodeLoginRequest,
		encodeLoginResponse,
	))

	return r
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UserReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeLoginResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
