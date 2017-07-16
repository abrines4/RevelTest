package controllers

import (
	"github.com/revel/revel"
	"github.com/Booking/app/models"
	"strconv"
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

	c.Session["user"] = strconv.FormatInt(auth_user.Id, 10)
	return c.Redirect(App.Home)
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

func (c App) Home() revel.Result {
	user_id, _ := strconv.ParseInt(c.Session["user"], 10, 64)
	revel.INFO.Println("HOME", user_id)
	user, err := UserCtrl(c).GetById(user_id)
	if err != nil || user.Id == 0 {
		revel.ERROR.Fatal("What even is this", err)
	}

	return c.Render(user)
}