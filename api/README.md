# API

Using go-swagger this time, as it seems more modern, more modular. This is an effort to ensure we can re-gen the code constantly without having to be careful about it.

## Generating the code
Go to the gen directory and run `make` 

Generating using [go-swagger](https://github.com/go-swagger/go-swagger)
[example repo](https://github.com/go-openapi/kvstore/blob/master/runtime.go)


## Executing the server
`go run ./cmd/main.go --scheme http`
