# transform-image-batch

Enablement Bootcamp for Gopherizing #2 の課題提出用リポジトリです。

## 制作物
コマンドラインでGCSのバケット名、オブジェクト名を渡すと
DBに保存済みの変換設定で画像変換する


```text
main.go // 変換実行ファイル
main_test.go // インテグレーションテストファイル（未作成分）
image.go //　画像変換処理ファイル
gcs
  gcs.go  // GCSのクライアントファイル
  gcs_test.go // gcs.goのユニットテストファイル（未作成）
  docker-compose.yml ローカルでGSSのモックを立ち上げるファイル
db
  db.go   // DBに接続するファイル
  client.go // GCSのクライアント作成ファイル
  db_test.go //db.goのユニットテストファイル（未作成）
  docker-compose.yml // ローカルでDBを立ち上げるdocker-composeファイル
  tmp 
    download // ダウンロードした画像を保存するディレクトリ
    transform // 変換後の画像を保存するディレクトリ
```

## 前提
- GCPアカウントでのASD設定済み
- リサイズ時にアスペクト比は考慮しない
- .envファイルにDB情報を記載してroot直下においてください

```
DB_USER=hogehoge
DB_PASS=password
MYSQL_HOST=xx.xx.xx.xx
PORT=xxxx
DBNAME=xxxx
```

※docker-compose.ymlファイルでローカル実行する場合は以下のように設定してください。
```
DB_USER=root
DB_PASS=<docker-compose.ymlで設定したパスワード>
MYSQL_HOST=127.0.0.1
PORT=8000
DBNAME=batch
```

## 動かし方
`go run　. -b <バケット名> -0 <オブジェクト名>`  
例）`go run　. -b "sample-bucket" -o "download/sample.jpg'"

```bash                                                                             4s   05:22:47 
❯ go run . -b "transform-image-bucket" -o "download/sample.jpeg"
Bucket Name: transform-image-bucket
Download Object Name: download/sample.jpeg
Connected!
setting: {1 sample.jpeg png 0.5 0.5}
Blob download/sample.jpeg downloaded to local file ./gcs/tmp/download/sample.jpeg
Image resized successfully!
/Users/kamadakohei/transform-image-batchBlob upload/20231023052422_formatted_resized_sample.jpeg uploaded.
Image uploaded successfully!
```


## 工夫した点
- 1日Go言語を学習した後、要件を満たすように最低限実装できた
  - Tour of go学習後、goのtutorialで今回の実装に必要な箇所だけ写経して実装した  
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
