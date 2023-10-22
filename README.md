# transform-image-batch

Enablement Bootcamp for Gopherizing #2 の課題提出用リポジトリです。

## 制作物
コマンドラインでGCSのバケット名、オブジェクト名（複数可）を渡すと
DBに保存済みの変換設定で画像変換する


```text
main.go // 変換実行ファイル
main_test.go // インテグレーションテストファイル
image.go //　画像変換用ファイル
gcs
  gcs.go  // GCSのクライアントファイル
  gcs_test.go // gcs.goのユニットテストファイル
  docker-compose.yml ローカルでGSSのモックを立ち上げるファイル
db
  db.go   // DBに接続するファイル
  db_test.go //db.goのユニットテストファイル
  docker-compose.yml // ローカルでDBを立ち上げるdocker-composeファイル

```

## 前提
- GCPアカウントでの認証済み
- アスペクト比率を変えずにリサイズすることは考慮しない
- .envファイルにDB情報を記載してroot直下においてください

```
DB_USER=hogehoge
DB_PASS=password
MYSQL_HOST=xx.xx.xx.xx
PORT=xxxx
DBNAME=xxxx
```


## 動かし方
`go run　. -b <バケット名> -0 <オブジェクト名>`  
例）`go run　. -b "sample-bucket" -o "download/sample.jpg'"


## 参考にしたドキュメント
- [go tutorial](https://go.dev/doc/tutorial/)
- [tour of go](https://go.dev/tour/list)
