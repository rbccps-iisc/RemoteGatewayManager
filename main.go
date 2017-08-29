package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rakshitadmar/gwCfgServer/api"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/gateways", api.GetAllGateways).Methods("GET")
	r.HandleFunc("/api/gateways/{id}", api.GetGateway).Methods("GET")
	r.HandleFunc("/api/register", api.PostGateway).Methods("POST")

	http.ListenAndServe(":3000", r)
}
