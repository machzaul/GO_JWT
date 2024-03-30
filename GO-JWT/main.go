package main

import (
	"go-jwt/configs"
	"go-jwt/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	configs.ConnectDB()

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	log.Println("server running on port 8090")
	http.ListenAndServe(":8090", router)

}
