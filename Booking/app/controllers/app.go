package controllers

import (
	"github.com/revel/revel"
	"github.com/Booking/app/models"
)

type App struct {
	GorpController
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Hello(name string, pass string) revel.Result {

	c.Validation.Required(name).Message("A name is required!")
	c.Validation.Required(pass).Message("A password is required!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	//Create user
	user := models.User{
		Name: name,
		Password: pass,
	}

	user.Validate(c.Validation)
	if c.Validation.HasErrors() {
		//TODO: make this less bad
		c.Flash.Error("User is not valid.")
	} else {
		if err := c.Txn.Insert(&user); err != nil {
			c.Flash.Error("Error inserting User into Database!")
		} else {
			c.Flash.Success("User Created!")
		}
	}

	c.FlashParams()
	return c.Redirect(App.Index)
}
