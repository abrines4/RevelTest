package controllers

import (
	"github.com/Booking/app/models"
)

type UserCtrl struct {
	GorpController
}


func (c UserCtrl) Create(user *models.User) error {
	err := c.Txn.Insert(user)
	return err
}

func (c UserCtrl) GetByName(name string) (*models.User, error) {
	user := new(models.User)
	err := c.Txn.SelectOne(user,
		`SELECT * FROM User WHERE name = ?`, name)
	
	return user, err
}

func (c UserCtrl) GetById(id int64) (*models.User, error) {
	user := new(models.User)
	err := c.Txn.SelectOne(user,
		`SELECT * FROM User WHERE id = ?`, id)
	
	return user, err
}