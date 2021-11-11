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
		t.FailNow()
		return
	}

	// リクエスト情報をcontextに入れる
	ginContext.Request = req

	ctrl := ShopsController{}
	ctrl.DB = db
	ctrl.Index(ginContext)
	// bodyStr := "[\n {\n \"id\": 11,\n \"name\": \"あっとほーむCafe\",\n \"address\": \"tokyo\",\n \"tel\": \"0469022323\",\n \"content\": \"fbsfbsfbsgmngsmgsnfbnrsbnsrjnsgnsg\",\n \"created_at\": \"2021-10-27T07:19:22Z\",\n \"updated_at\": \"2021-10-28T06:11:26Z\"\n },\n {\n \"id\": 13,\n \"name\": \"hogehoge\",\n \"address\": \"tokyo\",\n \"tel\": \"0469022323\",\n \"content\": \"fbsfbsfbsgmngsmgsnfbnrsbnsrjnsgnsg\",\n \"created_at\": \"2021-10-27T07:21:30Z\",\n \"updated_at\": \"2021-10-27T07:21:30Z\"\n },\n {\n \"id\": 14,\n \"name\": \"更新店舗\",\n \"address\": \"更新住所\",\n \"tel\": \"08978978\",\n \"content\": \"テスト\",\n \"created_at\": \"2021-10-27T08:56:21Z\",\n \"updated_at\": \"2021-11-05T04:00:08Z\"\n },\n {\n \"id\": 17,\n \"name\": \"アマリリス\",\n \"address\": null,\n \"tel\": null,\n \"content\": null,\n \"created_at\": \"2021-10-28T06:14:11Z\",\n \"updated_at\": \"2021-10-28T06:14:11Z\"\n },\n {\n \"id\": 18,\n \"name\": \"test cafe\",\n \"address\": null,\n \"tel\": null,\n \"content\": null,\n \"created_at\": \"2021-10-29T08:13:52Z\",\n \"updated_at\": \"2021-10-29T08:13:52Z\"\n }\n]"

	// 結果の確認
	// assert.Equal(t, bodyStr, w.Result().Body)
	assert.Equal(t, 200, w.Code)
}

// テストデータの用意
var post_tests = []struct {
	body string
}{
	// ここにデータを書いていく
	{"{\n \"shop\": {\n \"name\": \"cafe ketos1\",\n \"address\": \"四谷\",\n \"tel\": \"08978971\"\n },\n \"shop_genre_ids\":[\n 1,2\n ]\n}"},
	{"{\n \"shop\": {\n \"name\": \"cafe ketos2\",\n \"address\": \"新宿\",\n \"tel\": \"08978972\"\n },\n \"shop_genre_ids\":[\n 1,2\n ]\n}"},
	{"{\n \"shop\": {\n \"name\": \"cafe ketos3\",\n \"address\": \"渋谷\",\n \"tel\": \"08978973\"\n },\n \"shop_genre_ids\":[\n 1,2\n ]\n}"},
}

func TestCreate(t *testing.T) {
	db := db.Connect()
	for _, tt := range post_tests {
		ctrl := ShopsController{}
		ctrl.DB = db
		w := httptest.NewRecorder()

		// gin contextの生成
		ginContext, _ := gin.CreateTestContext(w)
		body := bytes.NewBufferString(tt.body)
		ginContext.Request, _ = http.NewRequest("POST", "/shops", body)

		ctrl.Create(ginContext)

		// 結果の確認
		assert.Equal(t, w.Code, 201)
	}
}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}
