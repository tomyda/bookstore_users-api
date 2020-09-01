package usersdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" //this is a justification comment :)
)

const (
	mysqlUsersUsername = "MY_SQL_USERS_USERNAME"
	mysqlUsersPassword = "MY_SQL_USERS_PASSWORD"
	mysqlUsersHost     = "MY_SQL_USERS_HOST"
	mysqlUsersSchema   = "MY_SQL_USERS_SCHEMA"
)

//Esto no lo use porque no supe como configurarlo.
//La idea era configurar la contraseñas y los datos ocultas desde el codigo.
var (
	//Client is a variable
	Client   *sql.DB
	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)
)

func init() {

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root", "12345678", "127.0.0.1:3306", "users_db",
	)
	var err error
	Client, err = sql.Open("mysql", datasourceName)

	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("DB succesfully configured")
}
