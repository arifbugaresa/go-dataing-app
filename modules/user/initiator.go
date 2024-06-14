package user

import (
	"github.com/gorilla/mux"
	"go-dating-app/databases"
)

func Initiate(r *mux.Router) {
	repo := NewRepository(databases.DBConnection)
	svc := NewService(repo)
	endpoint := Endpoints{
		LoginEndpoint: MakeLoginEndpoint(svc),
	}

	_ = NewHTTPHandler(endpoint, r)
}
