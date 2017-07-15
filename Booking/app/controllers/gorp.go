package controllers

import (
	"github.com/go-gorp/gorp"
	"database/sql"
	"github.com/revel/revel"
	"fmt"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/Booking/app/models"
)

//Almost entirely based on https://rclayton.silvrback.com/revel-gorp-and-mysql

var (
	Dbm *gorp.DbMap
)

type GorpController struct {
	*revel.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() revel.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}

	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}

	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}

	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}

	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}

	c.Txn = nil
	return nil
}


func getParamString(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Could not find parameter: " + param)
		} else {
			return defaultValue
		}
	}
	return p
}

func getConnectionString() string {
	host := getParamString("db.host", "")
	port := getParamString("db.port", "3306")
	user := getParamString("db.user", "")
	pass := getParamString("db.password", "")
	dbname := getParamString("db.name", "Booking")
	protocol := getParamString("db.protocol", "tcp")
	dbargs := getParamString("dbargs", " ")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}

	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s", 
		user, pass, protocol, host, port, dbname, dbargs)
}

func defineUserTable(dbm *gorp.DbMap){
	// set "id" as primary key and autoincrement
	t := dbm.AddTable(models.User{}).SetKeys(true, "id") 
	// e.g. VARCHAR(30)
	t.ColMap("name").SetMaxSize(30)
}

var InitDb func() = func() {
	revel.INFO.Println("INITDB")
	connectionString := getConnectionString()
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		revel.ERROR.Fatal(err)
	}

	Dbm = &gorp.DbMap {
		Db: db,
		Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"},
	}

	defineUserTable(Dbm)
	if err := Dbm.CreateTablesIfNotExists(); err != nil {
		revel.ERROR.Fatal(err)
	}
}
