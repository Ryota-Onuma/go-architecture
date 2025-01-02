# オニオンアーキテクチャ


Jeffrey Palermo氏の[The Onion Architecture : part 3](https://jeffreypalermo.com/2008/08/the-onion-architecture-part-3/) を参考にした。

依存の方向を以下としている。
Domain Model ← Repository/Domain services ← Application services(Use cases) ← Presentation/Infrastructure services
- 依存の方向さえ守っていれば良いので、Domain Modelとかを直接オニオンの外側の層とかから呼び出すとかはOK

層の間のやりとりはInterfaceを介す。


サーバー起動

```shell
go run cmd/main.gp
```

記事取得

```shell
curl -X POST http://localhost:3000/articles
```

記事作成

```shell
curl -X POST http://localhost:3000/article \
     -d "title=記事タイトルです" \
     -d "body=あいうえおかきくけこさしすせそ"
```

