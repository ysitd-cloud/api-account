// No Rights Reserved. Use of this API source code is governed by
// CC0 1.0 Universal (CC0 1.0) Public Domain Dedication.
// You can find a copy of license on https://creativecommons.org/publicdomain/zero/1.0/legalcode

package account

import (
	"context"

	"google.golang.org/grpc"
)

//go:generate protoc -I . -I vendor --gogo_out=plugins=grpc:. ./service.proto

type Client struct {
	Endpoint string
}

func NewClient(endpoint string) *Client {
	return &Client{
		Endpoint: endpoint,
	}
}

func (c *Client) connect(ctx context.Context) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx, c.Endpoint, grpc.WithInsecure())
}

func (c *Client) ValidateUserPassword(ctx context.Context, in *ValidateUserRequest, opts ...grpc.CallOption) (out *ValidateUserReply, err error) {
	cc, err := c.connect(ctx)
	if err != nil {
		return
	}

	defer cc.Close()

	out = new(ValidateUserReply)
	if err = cc.Invoke(ctx, "/account.Account/validateUserPassword", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (out *GetUserInfoReply, err error) {
	cc, err := c.connect(ctx)
	if err != nil {
		return
	}

	defer cc.Close()

	out = new(GetUserInfoReply)
	if err = cc.Invoke(ctx, "/account.Account/getUserInfo", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) GetTokenInfo(ctx context.Context, in *GetTokenInfoRequest, opts ...grpc.CallOption) (out *GetTokenInfoReply, err error) {
	cc, err := c.connect(ctx)
	if err != nil {
		return
	}

	defer cc.Close()

	out = new(GetTokenInfoReply)
	if err = cc.Invoke(ctx, "/account.Account/getTokenInfo", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}
