package main

import (
	"fmt"

	"github.com/akuruwasaki/concafe_map_mock/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World")
	})
	db := db.Connect()
	defer db.Close()

	err := db.Ping()

	if err != nil {
		fmt.Println(err)

		fmt.Println("データベース接続失敗")
	} else {
		fmt.Println("データベース接続成功")
	}
}
