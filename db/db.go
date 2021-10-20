package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {

	user := os.Getenv("DB_USER")
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("HOST_DB_PORT")
	database_name := os.Getenv("MYSQL_DATABASE")

	dbconf := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4"
	db, err := sql.Open("mysql", dbconf)
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
