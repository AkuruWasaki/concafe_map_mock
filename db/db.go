package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {

	// TODO: envファイルから情報を取得するようにする
	user := "root"
	password := "root"
	host := "localhost"
	port := "3306"
	database_name := "concafe_map_db"

	dbconf := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4"
	db, err := sql.Open("mysql", dbconf)
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
