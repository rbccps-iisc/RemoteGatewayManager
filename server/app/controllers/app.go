package controllers

import (
	//"github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Hello(myName string) revel.Result {
	return c.Render(myName)
}

func (c App) Login() revel.Result {

	return c.Render()
}

func (c App) Auth() revel.Result {

	uname := c.Params.Form.Get("username")
	pwd := c.Params.Form.Get("password")

	if uname == user.Username && pwd == user.Password {

		c.Session["username"] = uname
		hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
		if err == nil {
			c.Session["password"] = string(hash)
		}
		return c.Redirect(Inventory.Gateways)

	} else {
		c.Flash.Error("Login Failed")

		return c.Redirect(App.Login)
	}
}

func (c App) Logout() revel.Result {

	delete(c.Session, "username")
	delete(c.Session, "password")

	return c.Redirect(App.Login)
}
