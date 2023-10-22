# transform-image-batch

Enablement Bootcamp for Gopherizing #2 の課題提出用リポジトリです。

## 制作物
コマンドラインでGCSのバケット名、オブジェクト名（複数可）を渡すと
DBに保存済みの変換設定で画像変換する


```text
main.go // 変換実行ファイル
gcs
  gcs.go  // GCSのクライアントファイル
  gcs_test.go // gcs.goのユニットテストファイル
  docker-compose.yml ローカルでGSSのモックを立ち上げるファイル
db
  db.go   // DBに接続するファイル
  db_test.go //db.goのユニットテストファイル
  docker-compose.yml // ローカルでDBを立ち上げるdocker-composeファイル
main_test.go // インテグレーションテストファイル
```

## 前提
- GCPアカウントでの認証済み
- アスペクト比率を変えずにリサイズすることは考慮しない


## 動かし方
`go run -b <バケット名> -p <オブジェクト名>`


## 参考にしたドキュメント
- [go tutorial](https://go.dev/doc/tutorial/)
- [tour of go](https://go.dev/tour/list)
