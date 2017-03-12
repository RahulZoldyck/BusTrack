package controllers

import (
	"github.com/revel/revel"
	//	"github.com/RahulZoldyck/Bustrack/app/models"
)

type App struct {
	*revel.Controller
}

func LoginAuth(userName, password string) bool {
	//TODO: should write auth calls
	return true
}
func (c App) Index() revel.Result {
	return c.Render()
}
func (c App) Hello(myname string) revel.Result {
	c.Validation.Required(myname).Message("YOur name is required!")
	c.Validation.MinSize(myname, 3).Message("Your name is not long enough")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	return c.Render(myname)
}

func (c App) Login() revel.Result {
	userName := c.Params.Form.Get("userName")
	password := c.Params.Form.Get("password")
	if userName == "" && password == "" {
		return c.Render()
	}
	//TODO: check how to distinguish between POST and GET request
	c.Validation.MinSize(userName, 2).Message("userName is too short")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Login)
	}
	if LoginAuth(userName, password) {
		return c.Redirect(App.Register, "ffhsdo")
	} else {
		return c.Render("Username or Password is wrong!")
	}
}

func (c App) Register() revel.Result {
	return c.Render()
}
