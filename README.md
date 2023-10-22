# transform-image-batch

Enablement Bootcamp for Gopherizing #2 の課題提出用リポジトリです。

## 制作物
コマンドラインでGCSのバケット名、オブジェクト名（複数可）を渡すと
DBに保存済みの変換設定で画像変換する


```text
main.go // 変換実行ファイル
main_test.go // インテグレーションテストファイル
image.go //　画像変換処理ファイル
gcs
  gcs.go  // GCSのクライアントファイル
  gcs_test.go // gcs.goのユニットテストファイル
  docker-compose.yml ローカルでGSSのモックを立ち上げるファイル
db
  db.go   // DBに接続するファイル
  client.go // GCSのクライアント作成ファイル
  db_test.go //db.goのユニットテストファイル
  docker-compose.yml // ローカルでDBを立ち上げるdocker-composeファイル
  tmp 
    download // ダウンロードした画像を保存するディレクトリ
    transform // 変換後の画像を保存するディレクトリ
```

## 前提
- GCPアカウントでの認証済み
- リサイズ時にアスペクト比は考慮しない
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


## 工夫した点
- 半日ほどでGo言語を学習した後、実装した
  - Tour of go学習後、今回の実装に必要な箇所だけ写経して実装した
    https://github.com/kamadakohei/go-tutorial
- テスト時にローカル実行できる用のdocker-compose.ymlを作成した（ただし、時間の関係でテストを書けていない）
- main.goにすべての処理を書かずに、ディレクトリを分けて処理を書いた
- エラー処理を書いた
- ログを出力するようにした

## できなかったこと
- テストを書く
- Go言語のベストなディレクトリ構成
- エラーハンドリングのベストプラクティスに沿った実装
- CLI化

## 参考にしたドキュメント
- [go tutorial](https://go.dev/doc/tutorial/)
- [tour of go](https://go.dev/tour/list)
