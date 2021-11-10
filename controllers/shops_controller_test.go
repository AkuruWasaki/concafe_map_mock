package controllers

import (
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

	// テストデータの作成
	// TODO: 余裕があったらテストDBを作成してここではデータ作成の処理を入れない
	// var shop models.Shop
	// shop.Name = "cafe ketos"
	// shop.Address = null.StringFrom("新宿区四谷")
	// shop.Tel = null.StringFrom("0001112222")
	// shop.Content = null.StringFrom("サンプルの喫茶店です。")
	// if err := shop.Insert(ginContext, db, boil.Infer()); err != nil {
	// 	// エラー処理を入れる
	// 	return
	// }

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
	assert.Equal(t, w.Body.String(), "")
}

func TestCreate(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}
