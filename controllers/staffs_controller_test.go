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

func setupStaffRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	s := r.Group("/staffs")
	{
		ctrl := StaffsController{}
		ctrl.DB = db
		s.GET("", ctrl.Index)
		s.POST("", ctrl.Create)
		s.PUT("/:id", ctrl.Update)
		s.DELETE("/:id", ctrl.Delete)
	}

	return r
}

func TestIndexStaff(t *testing.T) {
	db := db.Connect()
	r := setupStaffRouter(db)
	w := httptest.NewRecorder()
	// gin contextの生成
	ginContext, _ := gin.CreateTestContext(w)

	// リクエストの生成
	req, err := http.NewRequest("GET", "/staffs", nil)
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
var staffs = []struct {
	body   string
	status int
}{
	// ここにデータを書いていく
	// name 未入力 登録に失敗すること
	{"{\n \"name\": null,\n \"age\": 23,\n \"shop_id\": 500\n}", 400},
	// name 1文字入力 登録に成功すること
	{"{\n \"name\": \"従\",\n \"age\": 23,\n \"shop_id\": 500\n}", 201},
	// name 255文字入力 登録に成功すること
	{"{\n \"name\": \"従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員\",\n \"age\": 23,\n \"shop_id\": 500\n}", 201},
	// name 256文字入力 登録に失敗すること
	{"{\n \"name\": \"従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員名テスト従業員テ\",\n \"age\": 23,\n \"shop_id\": 500\n}", 400},

	// age 未入力 登録に失敗すること
	{"{\n \"name\": \"サンプル従業員\",\n \"age\": null,\n \"shop_id\": 500 \n}", 400},
	// age 入力 登録に成功すること
	{"{\n \"name\": \"従業員名テスト\",\n \"age\": 23,\n \"shop_id\": 500\n}", 201},

	// // shop_id 未入力 登録に失敗すること
	{"{\n \"name\": \"従業員名テスト\",\n \"age\": 23,\n \"shop_id\": null\n}", 400},
	// // shop_id 存在する店舗IDを入力 登録に成功すること
	{"{\n \"name\": \"従業員名テスト\",\n \"age\": 23,\n \"shop_id\": 500\n}", 201},
	// // shop_id 存在しない店舗IDを入力 登録に失敗すること
	{"{\n \"name\": \"従業員名テスト\",\n \"age\": 1,\n \"shop_id\": 5000\n}", 400},
	// // shop_id 文字列を入力 登録に失敗すること
	// {"{\n \"name\": \"従業員名テスト\",\n \"age\": 1,\n \"shop_id\": \"500\"\n}", 400},
}

func TestCreateStaff(t *testing.T) {
	db := db.Connect()
	r := setupStaffRouter(db)
	for _, tt := range staffs {
		w := httptest.NewRecorder()
		// gin contextの生成
		ginContext, _ := gin.CreateTestContext(w)
		body := bytes.NewBufferString(tt.body)
		req, err := http.NewRequest("POST", "/staffs", body)
		if err != nil {
			log.Printf("error: %v", err)
		}
		ginContext.Request = req
		r.ServeHTTP(w, ginContext.Request)

		// 結果の確認
		assert.Equal(t, w.Code, tt.status)
		log.Println(w.Body)
	}
}
