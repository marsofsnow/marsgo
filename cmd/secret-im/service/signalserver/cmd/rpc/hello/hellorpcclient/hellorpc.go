// Code generated by goctl. DO NOT EDIT!
// Source: hello.proto

//go:generate mockgen -destination ./hellorpc_mock.go -package hellorpcclient -source $GOFILE

package hellorpcclient

import (
	"context"

	"secret-im/service/signalserver/cmd/rpc/hello/hello_rpc"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	Request       = hello_rpc.Request
	Response      = hello_rpc.Response
	IdReq         = hello_rpc.IdReq
	UserInfoReply = hello_rpc.UserInfoReply

	HelloRpc interface {
		Ping(ctx context.Context, in *Request) (*Response, error)
		GetUser(ctx context.Context, in *IdReq) (*UserInfoReply, error)
	}

	defaultHelloRpc struct {
		cli zrpc.Client
	}
)

func NewHelloRpc(cli zrpc.Client) HelloRpc {
	return &defaultHelloRpc{
		cli: cli,
	}
}

func (m *defaultHelloRpc) Ping(ctx context.Context, in *Request) (*Response, error) {
	client := hello_rpc.NewHelloRpcClient(m.cli.Conn())
	return client.Ping(ctx, in)
}

func (m *defaultHelloRpc) GetUser(ctx context.Context, in *IdReq) (*UserInfoReply, error) {
	client := hello_rpc.NewHelloRpcClient(m.cli.Conn())
	return client.GetUser(ctx, in)
}
