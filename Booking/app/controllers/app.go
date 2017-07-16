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

func (c App) Login(user models.User) revel.Result {
	auth_user, err := UserCtrl(c).GetByName(user.Name)
	if err != nil || auth_user.Id == 0 {
		c.Flash.Error("Incorrect Login")
		return c.Redirect(App.Index)
	}

	if user.Password != auth_user.Password {
		c.Flash.Error("Incorrect Login")
		return c.Redirect(App.Index)	
	}

	c.Flash.Success("Valid Credentials!")
	return c.Redirect(App.Index)
}

func (c App) Register() revel.Result {
	return c.Render()
}

func (c App) SaveUser(new_user models.User) revel.Result {
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
