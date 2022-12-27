GolangとGolangのWebアプリケーションフレームワークEchoの練習を兼ねて作った、JSON-based SQL injectionの検証用の簡易Webアプリケーション……のはずが、結局SQLインジェクションなら大体なんでもござれになりました。

## 起動方法

```
$ git clone https://github.com/shinobe179/json-object-sqli-sandbox.git
$ cd json-object-sqli-sandbox
$ docker-compose up -d
```

## 使い方

- `name` というクエリパラメータを与えると、Webアプリケーションはそれを元にMySQLへクエリします。
- クエリは文字列結合で構築されるので、SQLインジェクションが可能です。

### チャレンジ

- `someservice.users` テーブルに格納されているレコードを全件（4件）まとめて表示させよう！
- `secret.flags` テーブルに格納されているflagを表示させよう！

## 本当に検証したかったこと

JSONオブジェクトを使った恒等式によるSQLインジェクション。Echo（urlパッケージ？）はクエリパラメータ`;`が含まれているとパースに失敗するらしく、`%3B`に置き換える必要がある。

`http://localhost:1323/?name=%27%20OR%20JSON_EXTRACT(%27{%22key%22:%22value%22}%27,%20%27$.key%27)%20=%20%27value%27%3B%20--+`
