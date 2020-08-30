package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //this is a justification comment :)
)

var (
	//Client is a variable
	Client *sql.DB
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root",
		"root", //contraseña
		"127.0.0.1:3306",
		"users_db",
	)
	var err error
	Client, err := sql.Open("mysql", datasourceName)

	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("DB succesfully configured")
}
