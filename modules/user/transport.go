package user

import "github.com/gorilla/mux"

func Initiate(r *mux.Router) {
	r.HandleFunc("/login", Login).Methods("POST")
}
