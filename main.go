package main

import (
	"fmt"

	"github.com/akuruwasaki/concafe_map_mock/db"
	"github.com/akuruwasaki/concafe_map_mock/server"
)

func main() {
	db := db.Connect()
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("データベース接続失敗")
		panic(err)
	}

	server.Init(db)
}
