# オニオンアーキテクチャ


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

