package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	Client *sql.DB
)

const (
	mysql_users_username = "root"
	mysql_users_password = ""
	mysql_users_host     = "127.0.0.1:3306"
	mysql_users_schema   = "users_db"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", mysql_users_username, mysql_users_password, mysql_users_host, mysql_users_schema)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("Database successfully configured")
}
