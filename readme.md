
## コマンド
モジュールまとめてダウンロード
```shell
go mod tidy
```

## コンパイル
コンパイル後のファイル名には.pbがつく。

```shell

protoc -I . --go_out=. proto/emplayee.proto protp/date.proto

or
protoc -I . --go_out=. proto/*.proto

```

## grpcのserverとclientのオブションファイルを生成する
```shell

protoc -I. --go_out=. --go-grpc_out=. protoc/*.proto
```

```shell
--cpp_out=OUT_DIR
--csharp_out=OUT_DIR
--java_out=OUT_DIR
--js_out=OUT_DIR
--python_out=OUT_DIR

 ...etc
```


## http/1.1の課題
- リクエストの多重化
1リクエストに対して1レスポンスの制約。
- プロトコルオーバーヘッド
Cookieやトークンなどの毎回リクエストヘッダに付与してリクエストするため、オーバーヘッドが大きくなる

## http/2の特徴
- ストリームという概念を導入
- サーバープッシュ
クライアントからリクエスト無しにサーバーからデータを送信できる