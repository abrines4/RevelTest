package controllers

import (
	"github.com/Booking/app/models"
	"github.com/revel/revel"
	"encoding/json"
)

type UserCtrl struct {
	GorpController
}

func (c UserCtrl) parseUser() (models.User, error) {
    user := models.User{}
    err := json.NewDecoder(c.Request.Body).Decode(&user)
    return user, err
}

func (c UserCtrl) Add() revel.Result {
	if user, err := c.parseUser(); err != nil {
		return c.RenderText("Unable to parse the User from JSON.")
	} else {
		//Validate the model
		user.Validate(c.Validation)
		if c.Validation.HasErrors() {
			//TODO: make this less bad
			return c.RenderText("User is not valid.")
		} else {
			if err := c.Txn.Insert(&user); err != nil {
				return c.RenderText("Error inserting User into Database!")
			} else {
				return c.RenderJSON(user)
			}
		}
	}
}

func (c UserCtrl) Get(id int64) revel.Result {
	user := new(models.User)
	err := c.Txn.SelectOne(user,
		`SELECT * FROM User WHERE id = ?`, id)
	if err != nil {
		return c.RenderText("Error. No user with id found")
	}

	return c.RenderJSON(user)
}

func (c UserCtrl) Update(id int64) revel.Result {
	user, err := c.parseUser()
	if err != nil { 
		return c.RenderText("Unable to parse the User from JSON.")
	}

	user.Id = id
	success, err := c.Txn.Update(&user)
	if err != nil || success == 0 {
		return c.RenderText("Unable to update user.")
	}

	return c.RenderText("Updated %v", id)
}

func (c UserCtrl) Delete(id int64) revel.Result {
	success, err := c.Txn.Delete(&models.User{Id: id})
	if err != nil || success == 0 {
        return c.RenderText("Failed to remove User")
    }
    return c.RenderText("Deleted %v", id)
}