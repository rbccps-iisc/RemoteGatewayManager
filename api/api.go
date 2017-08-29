package api

import (
	"net/http"

	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rakshitadmar/gwCfgServer/db"
)

func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}

// GetAllGateways returns a list of all database gateways to the response.
func GetAllGateways(w http.ResponseWriter, req *http.Request) {
	rs, err := db.GetAll()
	if err != nil {
		handleError(err, "Failed to load database gateway: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to load marshal data: %v", w)
		return
	}

	w.Write(bs)
}

// GeGateway returns a single database gateway matching given MAC parameter.
func GetGateway(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	mac := vars["mac"]

	rs, err := db.GetOne(mac)
	if err != nil {
		handleError(err, "Failed to read database: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to marshal data: %v", w)
		return
	}

	w.Write(bs)
}

// PostGateway saves an gateway (form data) into the database.
func PostGateway(w http.ResponseWriter, req *http.Request) {
	MAC := req.FormValue("mac")
	IP := req.FormValue("ip")

	gateway := db.Gateway{MAC: MAC, IP: IP}

	if err := db.Save(gateway); err != nil {
		handleError(err, "Failed to save data: %v", w)
		return
	}

	w.Write([]byte("OK"))
}

// DeleteGateway removes a single gateway (identified by parameter) from the database.
func DeleteGateway(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	mac := vars["mac"]

	if err := db.Remove(mac); err != nil {
		handleError(err, "Failed to remove gateway: %v", w)
		return
	}

	w.Write([]byte("OK"))
}
