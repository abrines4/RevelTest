package models

import (
	"github.com/revel/revel"
	"regexp"
)

type User struct {
	Id	int64	`db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
}

func (user *User) Validate(validation *revel.Validation) {
	validation.Check(user.Name,
		revel.ValidRequired(),
		revel.ValidMaxSize(30),
	)

	validation.Check(user.Password,
		revel.ValidRequired(),
		revel.ValidMatch(regexp.MustCompile("^[[:graph:]]{8,}$")),
	)
}