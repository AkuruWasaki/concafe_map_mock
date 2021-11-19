package controllers

import (
	"bytes"
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akuruwasaki/concafe_map_mock/db"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func setUpRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	s := r.Group("/shops")
	{
		ctrl := ShopsController{}
		ctrl.DB = db
		s.GET("", ctrl.Index)
		s.POST("", ctrl.Create)
		s.PUT("/:id", ctrl.Update)
		s.DELETE("/:id", ctrl.Delete)
	}

	return r
}

func TestIndexShop(t *testing.T) {
	db := db.Connect()
	r := setUpRouter(db)
	w := httptest.NewRecorder()
	// gin contextの生成
	ginContext, _ := gin.CreateTestContext(w)

	// リクエストの生成
	req, err := http.NewRequest("GET", "/shops", nil)
	if err != nil {
		log.Printf("error: %v", err)
	}

	// リクエスト情報をcontextに入れる
	ginContext.Request = req

	r.ServeHTTP(w, ginContext.Request)

	// 結果の確認
	// assert.Equal(t, bodyStr, w.Result().Body)
	assert.Equal(t, 200, w.Code)
}

// テストデータの用意
var post_tests = []struct {
	body   string
	status int
}{
	// ここにデータを書いていく
	// WIP: バリデーション実装する

	// name 未入力 登録に失敗すること
	{"{\n \"shop\": {\n \"name\": \"\",\n \"address\": \"四谷\",\n \"tel\": \"08978971\"\n },\n \"shop_genre_ids\":[\n 1,2\n ]\n}", 400},
	// name 空白スペース入力 登録に失敗すること
	// TODO: white spaceに関する対応はvalidatorに標準搭載されていないため, 自分で実装する
	// {"{\n \"shop\": {\n \"name\": \" \",\n \"address\": \"更新住所\",\n \"tel\": \"a\"\n },\n \"shop_genre_ids\":[\n 1,2\n ]\n}", 400},
	// name 255文字 登録に成功すること
	{"{\n \"shop\": {\n \"name\": \"cafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-\",\n \"address\": \"更新住所\",\n \"tel\": \"08978978\"\n },\n \"shop_genre_ids\":[\n 1,2\n ]\n}", 201},
	// name 256文字 登録に失敗すること
	{"{\n \"shop\": {\n \"name\": \"cafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-ketoscafe-k\",\n \"address\": \"四谷\",\n \"tel\": \"08978971\"\n },\n \"shop_genre_ids\":[\n 1,2\n ]\n}", 400},

	// address
	// 未入力 登録に成功すること
	{"{\n \"shop\": {\n \"name\": \"cafe-ketos\",\n \"address\": \"\",\n \"tel\": \"08978978\"\n },\n \"shop_genre_ids\":[\n 1,2\n ]\n}", 201},
	// 255文字 登録に成功すること
	{"{\n \"shop\": {\n \"name\": \"cafe-ketos\",\n \"address\": \"新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷\",\n \"tel\": \"08978978\"\n },\n \"shop_genre_ids\":[\n 1,2\n ]\n}", 201},
	// 256文字 登録に失敗すること
	{"{\n \"shop\": {\n \"name\": \"cafe-ketos\",\n \"address\": \"新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新宿区四谷新\",\n \"tel\": \"08978978\"\n },\n \"shop_genre_ids\":[\n 1,2\n ]\n}", 400},

	// tel
	// 未入力 登録に成功すること
	{"{\n \"shop\": {\n \"name\": \"cafe-ketos\",\n \"address\": \"新宿四谷\",\n \"tel\": \"\"\n },\n \"shop_genre_ids\":[\n 1,2\n ]\n}", 201},
	// 45文字 登録に成功すること
	{"{\n \"shop\": {\n \"name\": \"cafe-ketos\",\n \"address\": \"新宿四谷\",\n \"tel\": \"080657762831312080657762831312080657762831312\"\n },\n \"shop_genre_ids\":[\n 1,2\n ]\n}", 201},
	// 46文字 登録に失敗すること
	{"{\n \"shop\": {\n \"name\": \"cafe-ketos\",\n \"address\": \"新宿四谷\",\n \"tel\": \"0806577628313120806577628313120806577628313121\"\n },\n \"shop_genre_ids\":[\n 1,2\n ]\n}", 400},

	// content
	// 未入力 登録に成功すること
	{"{\n \"shop\": {\n \"name\": \"更新店舗\",\n \"address\": \"\",\n \"tel\": \"08978978\",\n \"content\": \"\"\n },\n \"shop_genre_ids\": [\n 1,\n 2\n ]\n}", 201},
	// 1000文字 登録に成功すること
	{"{\n \"shop\": {\n \"name\": \"更新店舗\",\n \"address\": \"\",\n \"tel\": \"08978978\",\n \"content\": \"これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェ\"\n },\n \"shop_genre_ids\": [\n 1,\n 2\n ]\n}", 201},
	// 1001文字 登録に失敗すること
	{"{\n \"shop\": {\n \"name\": \"更新店舗\",\n \"address\": \"\",\n \"tel\": \"08978978\",\n \"content\": \"これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの店舗情報です。これはカフェの\"\n },\n \"shop_genre_ids\": [\n 1,\n 2\n ]\n}", 400},

	// shopに紐づくshop_genre_idsがない場合、登録に成功すること
	{"{\n \"shop\": {\n \"name\": \"cafe ketos2\",\n \"address\": \"新宿\",\n \"tel\": \"08978972\"\n }}", 201},

	// shop_genre_idsの値がstring型だった時に登録に失敗すること
	{"{\n \"shop\": {\n \"name\": \"更新店舗\",\n \"address\": \"更新住所\",\n \"tel\": \"08978978\"\n },\n \"shop_genre_ids\":[\n \"1\", \"3\"\n ]\n}", 400},
}

func TestCreateShop(t *testing.T) {
	db := db.Connect()
	r := setUpRouter(db)
	for _, tt := range post_tests {
		w := httptest.NewRecorder()
		// gin contextの生成
		ginContext, _ := gin.CreateTestContext(w)
		body := bytes.NewBufferString(tt.body)
		req, err := http.NewRequest("POST", "/shops", body)
		if err != nil {
			log.Printf("error: %v", err)
		}
		ginContext.Request = req
		r.ServeHTTP(w, ginContext.Request)

		// 結果の確認
		assert.Equal(t, w.Code, tt.status)
	}
}

func TestUpdateShop(t *testing.T) {
	db := db.Connect()
	r := setUpRouter(db)
	w := httptest.NewRecorder()

	// gin contextの生成
	ginContext, _ := gin.CreateTestContext(w)
	body := bytes.NewBufferString("{\n \"name\": \"更新店舗\",\n \"address\": \"更新住所\",\n \"tel\": \"08978978\",\n \"content\": \"aaaa\"\n}")
	req, err := http.NewRequest("PUT", "/shops/492", body) // idはDBにあるデータにしておく。
	if err != nil {
		log.Printf("error: %v", err)
	}
	ginContext.Request = req

	r.ServeHTTP(w, ginContext.Request)

	// 結果の確認
	assert.Equal(t, w.Code, 200)
}

func TestDeleteShop(t *testing.T) {
	db := db.Connect()
	r := setUpRouter(db)
	w := httptest.NewRecorder()

	// gin contextの生成
	ginContext, _ := gin.CreateTestContext(w)
	//
	req, err := http.NewRequest("DELETE", "/shops/491", nil)
	if err != nil {
		log.Printf("error: %v", err)
	}
	ginContext.Request = req

	r.ServeHTTP(w, ginContext.Request)

	// 結果の確認
	assert.Equal(t, ginContext.Writer.Status(), 200)
}
