package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akuruwasaki/concafe_map_mock/db"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func TestIndex(t *testing.T) {
	db := db.Connect()
	w := httptest.NewRecorder()
	// gin contextの生成
	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	// リクエストの生成
	req, err := http.NewRequest("GET", "/shops", nil)
	if err != nil {
		// エラー処理を入れる
		return
	}

	// リクエスト情報をcontextに入れる
	ginContext.Request = req

	ctrl := ShopsController{}
	ctrl.DB = db
	ctrl.Index(ginContext)

	// 結果の確認
	assert.Equal(t, 200, w.Code)
}

func TestCreate(t *testing.T) {
	db := db.Connect()
	ctrl := ShopsController{}
	ctrl.DB = db
	w := httptest.NewRecorder()
	// gin contextの生成
	ginContext, _ := gin.CreateTestContext(w)
	body := bytes.NewBufferString("{\n \"shop\": {\n \"name\": \"cafe ketos\",\n \"address\": \"四谷\",\n \"tel\": \"08978978\"\n },\n \"shop_genre_ids\":[\n 1,2\n ]\n}")
	ginContext.Request, _ = http.NewRequest("POST", "/shops", body)
	ctrl.Create(ginContext)

	// 結果の確認
	assert.Equal(t, w.Code, 201)
}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}
