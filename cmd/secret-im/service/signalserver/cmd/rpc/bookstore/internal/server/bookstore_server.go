// Code generated by goctl. DO NOT EDIT!
// Source: bookstore.proto

package server

import (
	"context"

	"secret-im/service/signalserver/cmd/rpc/bookstore/bookstore"
	"secret-im/service/signalserver/cmd/rpc/bookstore/internal/logic"
	"secret-im/service/signalserver/cmd/rpc/bookstore/internal/svc"
)

type BookstoreServer struct {
	svcCtx *svc.ServiceContext
}

func NewBookstoreServer(svcCtx *svc.ServiceContext) *BookstoreServer {
	return &BookstoreServer{
		svcCtx: svcCtx,
	}
}

func (s *BookstoreServer) Ping(ctx context.Context, in *bookstore.Request) (*bookstore.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *BookstoreServer) Add(ctx context.Context, in *bookstore.AddReq) (*bookstore.AddResp, error) {
	l := logic.NewAddLogic(ctx, s.svcCtx)
	return l.Add(in)
}

func (s *BookstoreServer) Check(ctx context.Context, in *bookstore.CheckReq) (*bookstore.CheckResp, error) {
	l := logic.NewCheckLogic(ctx, s.svcCtx)
	return l.Check(in)
}
