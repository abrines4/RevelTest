package controllers

import "github.com/revel/revel"

func init() {
	revel.INFO.Println("CONTROLLER INIT")
	revel.OnAppStart(InitDb)
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}