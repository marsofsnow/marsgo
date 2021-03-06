// Code generated by goctl. DO NOT EDIT!
// Source: bookstore.proto

//go:generate mockgen -destination ./bookstore_mock.go -package bookstoreclient -source $GOFILE

package bookstoreclient

import (
	"context"

	"secret-im/service/signalserver/cmd/rpc/bookstore/bookstore"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	CheckResp = bookstore.CheckResp
	Request   = bookstore.Request
	Response  = bookstore.Response
	AddReq    = bookstore.AddReq
	AddResp   = bookstore.AddResp
	CheckReq  = bookstore.CheckReq

	Bookstore interface {
		Ping(ctx context.Context, in *Request) (*Response, error)
		Add(ctx context.Context, in *AddReq) (*AddResp, error)
		Check(ctx context.Context, in *CheckReq) (*CheckResp, error)
	}

	defaultBookstore struct {
		cli zrpc.Client
	}
)

func NewBookstore(cli zrpc.Client) Bookstore {
	return &defaultBookstore{
		cli: cli,
	}
}

func (m *defaultBookstore) Ping(ctx context.Context, in *Request) (*Response, error) {
	client := bookstore.NewBookstoreClient(m.cli.Conn())
	return client.Ping(ctx, in)
}

func (m *defaultBookstore) Add(ctx context.Context, in *AddReq) (*AddResp, error) {
	client := bookstore.NewBookstoreClient(m.cli.Conn())
	return client.Add(ctx, in)
}

func (m *defaultBookstore) Check(ctx context.Context, in *CheckReq) (*CheckResp, error) {
	client := bookstore.NewBookstoreClient(m.cli.Conn())
	return client.Check(ctx, in)
}
