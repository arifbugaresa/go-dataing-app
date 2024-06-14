package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-dating-app/databases"
	"go-dating-app/modules/user"
	"log"
	"net/http"
)

func main() {
	// Initialize MySQL database
	databases.InitDatabase()

	// Create router
	r := mux.NewRouter()

	// Route
	user.Initiate(r)

	// Start server
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
