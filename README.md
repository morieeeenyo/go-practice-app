# go-practice-app
以下の目的で作ったアプリ
- goの学習
- クリーンアーキテクチャの考え方を採用したバックエンド開発の練習

# 技術スタック
- Go
- Next.js
- Typescript

# 環境構築手順
まず最初にGoおよびNodeのインストールを行います。
開発時のバージョンはGo1.20系、Node18.12系です。

## バックエンド
以下のコマンドを実行して必要なパッケージをインストールします。

```
cd backend
go install github.com/cosmtrek/air@latest
go mod download && go mod verify
```

次にテーブルの作成に必要なgolang-migrateと、ORMであるsqlboilerを動かすためのパッケージ、ホットリロードを使うためのairをインストールします。
```
go install github.com/volatiletech/sqlboiler@latest
go install github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql@latest
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
```

サーバー起動時は以下のコマンドを実行します。

```
make start-backend
```

次にローカル環境に `go_practice_app_development` という名前でデータベースを作成してください。
コマンドラインでもMySQLクライアントからでもOKです。

## フロントエンド
以下のコマンドを実行して必要なパッケージをインストールします。

```
cd frontend
npm install
```

サーバー起動時は以下のコマンドを実行します。

```
make start-frontend
```

localhost:3000にアクセスし、以下の画面が表示されていれば環境構築完了です。

<img width="1439" alt="go-practice-app" src="https://github.com/morieeeenyo/go-practice-app/assets/64336740/238330de-91f7-4177-b9ba-519a612fdb63">

# 動作イメージ
軽いローディング画面を挟んだのち、TODOリストが表示されます。
画面上でTODOの追加ができます。

https://github.com/morieeeenyo/go-practice-app/assets/64336740/81ce2463-175b-4ad6-a911-a9025f1952eb
