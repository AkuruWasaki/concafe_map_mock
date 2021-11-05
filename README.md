# concafe_map_mock

# 何をするためのアプリか
喫茶店や飲食店の店舗情報, スタッフ情報を簡易的に管理するためのアプリです。

# テーブル概要説明
- shops 店舗情報
- staffs 従業員
- shop_genres 店舗ジャンルマスタ
- shop_genre_relations 店舗ジャンル中間テーブル

# 環境構築手順
※ VScodeにVSCode Remote Containerが入っている場合じゃ1, 2をスキップ
1. VSCodeでDockerfileやdocker-compose.ymlがあるプロジェクトを開く
2. 拡張機能(Remote-Containers)をインストール
3. VSCodeの画面の一番左下の部分にアイコンが出てくるので、クリック
4. Reopen in Containerをクリック
5. コンテナ内でVSCodeが開けば成功。初回はビルドがあるので時間がかかります。
6. ビルドに成功したら
```
## 下記コマンドをコンテナ内の/appディレクトリ直下で叩いてください
go mod init concafe_map
go mod tidy
```

# これから追加したい要素
-  テストコード
-  クライアント側を書く(JS or Flutter)
- gRPCでの通信

# 不安なのでレビュー欲しい点
-  エラーハンドリングにまだまだ甘い部分があると思われる
-  変数名等の命名。特に構造体の名前の付け方
-  shops_controllerでmany to manyの関係がある中間テーブルのデータ作成、更新をする際の処理は適切かどうか。