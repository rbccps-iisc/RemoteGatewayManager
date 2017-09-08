package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rakshitadmar/gwCfgServer/db"
	"io/ioutil"
	"net/http"
	"net"
	"regexp"
)

type Gw struct {
	Ip  string `json:"ip"`
	Mac string `json:"mac"`
	Username string `json:"username"`
}

type RespMsg struct {
	Port string `json:"port"`

}


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

	// var data map[string]interface{}
	// err := json.Unmarshal([]byte(req), &data)

	var gw Gw
	
	rexp, _ := regexp.Compile("[\\d]+")

	iface,_ := net.Listen("tcp",":0")
	defer iface.Close()
	freePort := rexp.FindString(iface.Addr().String())
	
	fmt.Println(freePort)

	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(b, &gw)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	MAC := gw.Mac
	fmt.Println(MAC)
	IP := gw.Ip
	fmt.Println(IP)
	Uname := gw.Username
	fmt.Println(Uname)
	gateway := db.Gateway{MAC: MAC, IP: IP, Port: freePort, Username: Uname}

	if err := db.Save(gateway); err != nil {
		handleError(err, "Failed to save data: %v", w)
		return
	}
	
	respm := RespMsg{Port:freePort}
	responseBody,_ := json.Marshal(respm)
	w.Write(responseBody)
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
