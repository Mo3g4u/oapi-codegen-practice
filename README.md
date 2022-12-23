# oapi-codegen-practice

- [deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen)

## generate gen.go
```sh
$ oapi-codegen -config config.yaml petstore-expanded.yaml > gen/openapi.gen.go
```

## start server
```sh
$ go run main.go
```

