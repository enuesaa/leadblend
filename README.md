# leadblend
データ整理ツール

- JSON 形式のデータを登録して、それを分類していく
- データは SQLite に保持する

## モチベーション
- ゴールデンウィークを利用して1つアプリケーションを作りたくなった
- 実装途中

## Stack
- Sveltekit
- Go
- GraophQL

## Commands
```bash
leadblend # editor を起動
```

## 概念
- space ... アーカイブファイルと1対1
- planet
- island ... name, content, tags, resources
- comet ... trash or memo

## Epilogue
- 実質2週間でよくここまで作ったなという印象。満足度が高い
- graph-gophers/graphql-go を初めて採用
  - Subscription のプロトコル graphql-transport-ws に未対応な件につまづいた
  - Query や Mutation といった一般的なユースケースにおいては使い勝手良く感じた
- sveltekit については例えば key block の存在を知るなど勉強になった
