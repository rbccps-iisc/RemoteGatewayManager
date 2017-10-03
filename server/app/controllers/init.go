package controllers

import (
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type UserCred struct {
	Username string
	Password string
}

var user UserCred

// simple example or user auth
func checkSession(c *revel.Controller) revel.Result {

	if c.Session["username"] == user.Username {
		if err := bcrypt.CompareHashAndPassword([]byte(c.Session["password"]), []byte(user.Password)); err != nil {
			c.Flash.Error("Login Failed")
			return c.Redirect(App.Login)
		} else {
			return nil
		}

	} else if (c.Request.Header.Get("username") == user.Username) && (c.Request.Header.Get("password") == user.Password) {
		return nil
	} else {
		c.Flash.Error("Login Failed")
		return c.Redirect(App.Login)

	}
}

func doNothing(c *revel.Controller) revel.Result { return nil }

func init() {

	user = UserCred{Username: "admin", Password: "admin"}
	revel.InterceptFunc(checkSession, revel.BEFORE, &Inventory{})

}
