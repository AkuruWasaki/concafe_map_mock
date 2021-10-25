package main

import (
	"fmt"

	"github.com/akuruwasaki/concafe_map_mock/db"
	"github.com/akuruwasaki/concafe_map_mock/server"
)

func main() {
	db := db.Connect()
	defer db.Close()

	err := db.Ping()

	if err != nil {
		fmt.Println("データベース接続失敗")
	} else {
		fmt.Println("データベース接続成功")
	}

	server.Init()
}
