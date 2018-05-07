// Copyright 2018 Tony Yip. All rights reserved.
// Use of this source code is governed by MIT license.

package account

//go:generate protoc -I . --go_out=plugins=grpc:. ./service.proto

type Client = AccountClient
