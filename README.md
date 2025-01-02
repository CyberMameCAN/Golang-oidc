# Project Golang-oidc

## 使用技術

- [Auth0](https://auth0.com)
- [blue-print](https://github.com/Melkeydev/go-blueprint)
- [Golang](https://go.dev/) + [Gin](https://github.com/gin-gonic/gin)
- etc...

[OpenID Connect](https://financial.ctc-g.co.jp/itinfo/2402-oidc)の仕組みを理解するために、上記技術等を使わせていただいて、Golangでハンズオンした忘備録になります。  
ここのソースコードは[Auth0](https://auth0.com)のサンプルほぼそのまま使っています。

### 参考にしたWebサイト

- [GoでOpenID ConnectのClientを実装する（Auth0設定編）](https://times.hrbrain.co.jp/entry/go-openid-connect-auth0-setting)
- [GoでOpenID ConnectのClientを実装する（実装編）](https://times.hrbrain.co.jp/entry/go-openid-connect-implement)
- [【Go】Goのプロジェクトの雛形を作れるCLIツール「go-blueprint」を使ってみる](https://qiita.com/_ken_/items/316d682b43d61cc5e34b)


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## env

    cp env.sample .env

.envの初期値をよしなに変更する。

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
