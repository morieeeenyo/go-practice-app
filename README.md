# go-practice-app
以下の目的で作ったアプリ
- goの学習
- クリーンアーキテクチャの考え方を採用したバックエンド開発の練習

# 技術スタック
- Go
- Next.js
- Typescript

# 動作イメージ
軽いローディング画面を挟んだのち、TODOリストが表示され画面上でTODOの追加ができる

https://github.com/morieeeenyo/go-practice-app/assets/64336740/81ce2463-175b-4ad6-a911-a9025f1952eb

## 環境構築手順
### バックエンド
以下のコマンを実行して必要なパッケージをインストールします。
テーブルの作成に必要なgolang-migrateと、ORMであるsqlboilerを動かすためのパッケージ、ホットリロードを使うためのairをインストールします。

```
cd backend
go install github.com/volatiletech/sqlboiler@latest
go install github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql@latest
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
go install github.com/cosmtrek/air@latest
go mod download && go mod verify
```

サーバー起動時は以下のコマンドを実行します。

```
make start-backend
```

### フロントエンド
以下のコマンを実行して必要なパッケージをインストールします。

```
cd frontend
npm install
```

サーバー起動時は以下のコマンドを実行します。

```
make start-frontend
```