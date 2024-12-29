# Project Golang-oidc

## 使用技術

- [Auth0](https://auth0.com)
- [blue-print](https://github.com/Melkeydev/go-blueprint)
- [Golang](https://go.dev/) + [Gin](https://github.com/gin-gonic/gin)
- etc...

OpenID Connectの仕組みを理解するために、上記技術等を使わせていただいて、Golangでハンズオンした防備録になります。  
ここのソースコードは[Auth0](https://auth0.com)のサンプルほぼそのまま使っています。

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

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
