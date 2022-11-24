# Go To Do Application

## Description

東京工業大学 情報理工学院 情報工学系 システム設計演習 (2022 3Q)

講義ページ: https://cs-sysdes.github.io/lecture-note/index.html

評価方法: https://cs-sysdes.github.io/about.html

レポート課題: https://cs-sysdes.github.io/report.html

仕様: https://cs-sysdes.github.io/todolist.html

## Environments

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
