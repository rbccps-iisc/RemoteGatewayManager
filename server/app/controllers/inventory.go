package controllers

import (
	//"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"github.com/rraks/RemoteGatewayManager/server/app/db"

	"os/exec"

	"net"

	"regexp"
)

type Gw struct {
	Ip       string `json:"ip"`
	Mac      string `json:"mac"`
	Username string `json:"username"`
}

type RespMsg struct {
	Port string `json:"port"`
}

type Inventory struct {
	*revel.Controller
}

func (c Inventory) Gateways() revel.Result {

	var gws []db.Gateway
	gws, _ = db.GetAll()

	return c.Render(gws)
}

func (c Inventory) Launch() revel.Result {

	var gw db.Gateway
	user := c.Params.Query.Get("user")
	port := c.Params.Query.Get("port")
	cmd := "gnome-terminal -e  \"ssh -t manager@139.59.88.117 ssh " + user + "@localhost" + "-p " + port + "\""

}

// PostGateway saves an gateway (form data) into the database.
func (c Inventory) Register() revel.Result {

	// var data map[string]interface{}
	// err := json.Unmarshal([]byte(req), &data)

	var gw Gw

	rexp, _ := regexp.Compile("[\\d]+")

	iface, _ := net.Listen("tcp", ":0")
	defer iface.Close()
	freePort := rexp.FindString(iface.Addr().String())

	//b, err := ioutil.ReadAll(c.Request.Body)

	// if err != nil {
	// 	return c.RenderText("Can't read body")
	// }

	c.Params.BindJSON(&gw)

	MAC := gw.Mac
	fmt.Println(MAC)
	IP := gw.Ip
	fmt.Println(IP)
	Uname := gw.Username
	fmt.Println(Uname)
	gateway := db.Gateway{MAC: MAC, IP: IP, Port: freePort, Username: Uname}

	if _, err := db.Save(gateway); err != nil {

		return c.RenderText("Couldn't save to DB")
	}

	respm := RespMsg{Port: freePort}
	//responseBody, _ := json.Marshal(respm)
	return c.RenderJSON(respm)
}
