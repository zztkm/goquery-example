# goquery example

## 実行

```
go get -u
```

引数にページのURLを入れる
```
go run main.go https://fortune.yahoo.co.jp/12astro/20190922/scorpio.html
```

## Build

for windows
```
GOOS=windows GOARCH=amd64 go build -o bin/goquery-example -trimpath
```

## 参考

- [【Go】Webスクレイピングのやり方 - Qiita](https://qiita.com/koki_develop/items/dab4bcbb1df1271a17b6)
