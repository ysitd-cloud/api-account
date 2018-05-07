# Account API (In-Cluster)

API for in-cluster API communication

## License
No Rights Reserved.
Use of this API source code is governed by CC0 1.0 Universal (CC0 1.0) Public Domain Dedication.
You can find a copy of license on https://creativecommons.org/publicdomain/zero/1.0/legalcode

## Requirement for change

- [protoc](https://github.com/google/protobuf)
- [Language Plugin of grpc](https://github.com/grpc/grpc/blob/master/INSTALL.md)

## How about other language?

There are two option:

1. Generate api from the grpc proto file
2. Just connect with the in-cluster API gateway

## Golang
```bash
go get -u github.com/golang/protobuf/protoc-gen-go
go generate ./...
```
