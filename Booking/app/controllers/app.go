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

func (c App) Register() revel.Result {
	return c.Render()
}

func (c App) SaveUser(new_user models.User) revel.Result {
	revel.INFO.Println("Save", new_user.Name, new_user.Password)
	if user, _ := UserCtrl(c).GetByName(new_user.Name); user.Id > 0 {
		c.Flash.Error("That name is taken.")
		revel.INFO.Println("TAKEN", user.Id, " ", user.Name)
		c.FlashParams()
		return c.Redirect(App.Register)
	}

	if err := UserCtrl(c).Create(&new_user); err != nil {
		c.Flash.Error("There was an error saving the user", err)
		return c.Redirect(App.Register)
	}

	c.Flash.Success("User Created!")
	return c.Redirect(App.Index)
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
