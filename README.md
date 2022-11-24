# Go To Do Application

## Description

東京工業大学 情報理工学院 情報工学系 システム設計演習 (2022 3Q)

講義ページ: https://cs-sysdes.github.io/lecture-note/index.html

評価方法: https://cs-sysdes.github.io/about.html

レポート課題: https://cs-sysdes.github.io/report.html

仕様: https://cs-sysdes.github.io/todolist.html

### Local Environment

macOS 13.0.1 (arm64)

## Environments

- Go Gin

  ドキュメント(日本語): https://gin-gonic.com/ja/docs/examples/

- Air: https://github.com/cosmtrek/air

  本プロジェクトでは、本番環境へのデプロイを想定していないため、開発体験を向上させる目的から、Air を使用している。

- wait-for-it:

  GitHub: https://github.com/vishnubob/wait-for-it

  docker-compose.yaml にて、DB が起動するまで待機するために使用している。

  具体的には、以下の部分である。

  ```yaml
  depends_on:
    - db
  command:
    - wait-for-it.sh
    - db:3306
    - --timeout=60
    - --strict
    - --
    - air
    - -c
    - .air.toml
  ```

  MySQL の起動を待ってから、Air を起動するようにしている。

## How to setup

`docker-compose up -d`を行うだけでよい。

詳細説明

- フロントエンド

  http://localhost:3000 にアクセスすると、root にアクセスできる。

  Chrome 拡張機能 React Developer Tools をインストールすると、React のコンポーネントを確認することができる。

- バックエンド

  http://localhost:8000 にアクセスすると、root にアクセスできる。
  基本的には、フロントエンドからのリクエストに対して、レスポンスを返すだけの役割を持つ。

  講義で扱ったような、HTML を返すような処理は行わない。（純粋な REST API として振る舞う）

- データベース

  phpMyAdmin は導入していない。代わりに Table Plus でのアクセス方法を記す

  - Name: 好きな名前(例: todo-application)
  - Host: 0.0.0.0
  - Port: 3306
  - User: developer
  - Password: password

  以下のようなになっていれば、接続できる。

  ![image](public/table-plus.png)

## Directory Structure

```bash
.
├── README.md
├── docker
│   ├── development   <-- for local development
│   └── production    <-- for production
├── docker-compose.yaml
├── docs
│   ├── docker.md
│   └── git.md
├── server            <-- server side
│   ├── api
│   ├── config
│   ├── controllers
│   ├── db
│   ├── go.mod
│   ├── go.sum
│   ├── helpers
│   ├── main.go
│   ├── models
│   ├── router
│   └── tmp
└── web               <-- client side
    ├── README.md
    ├── next-env.d.ts
    ├── next.config.js
    ├── node_modules
    ├── package.json
    ├── pages
    ├── public
    ├── styles
    ├── tsconfig.json
    └── yarn.lock
```
