# Steps to setup your dev environment

- Install go
- Clone this repository
- Get gorilla/mux package
    `go get -u github.com/gorilla/mux`
- Get driver for postgres db
    `go get -u github.com/lib/pq`
- Get the Zap package
    `go get -u go.uber.org/zap`
- Run with the service with
    `go run src/main.go`
- If you want try the samples from `\http_client\samples.http` you need a vscode extension `REST Client`
