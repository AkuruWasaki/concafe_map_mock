# concafe_map_mock

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
