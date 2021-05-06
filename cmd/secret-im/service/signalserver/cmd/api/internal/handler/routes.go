// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	accounts "secret-im/service/signalserver/cmd/api/internal/handler/accounts"
	bookstore "secret-im/service/signalserver/cmd/api/internal/handler/bookstore"
	msgs "secret-im/service/signalserver/cmd/api/internal/handler/msgs"
	textsecret "secret-im/service/signalserver/cmd/api/internal/handler/textsecret"
	website "secret-im/service/signalserver/cmd/api/internal/handler/website"
	"secret-im/service/signalserver/cmd/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/index",
				Handler: IndexHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: website.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: website.LoginHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserCheck},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/user/info",
					Handler: website.UserInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/add",
				Handler: bookstore.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/check",
				Handler: bookstore.CheckHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/v1/accounts/sms/code/:number",
				Handler: accounts.GetSmsCodeHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckBasicAuth},
			[]rest.Route{
				{
					Method:  http.MethodPut,
					Path:    "/v1/accounts/code/code/:verificationCode",
					Handler: accounts.ConfirmCodeHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/v1/accounts/attributes",
					Handler: accounts.PutAttributesHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/devices",
					Handler: accounts.GetDevicesHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/v1/devices/:deviceId",
					Handler: accounts.DeleteDeviceHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/v1/devices/:verificationCode",
					Handler: accounts.PutDeviceHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/profile/:number",
					Handler: accounts.GetProfileHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/directory/:token",
					Handler: accounts.GetDirTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/directory/tokens",
					Handler: accounts.GetDirTokensHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/keys",
					Handler: accounts.GetKeysHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/keys/:number/:deviceId",
					Handler: accounts.GetDeviceKeyHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/v1/keys",
					Handler: accounts.PutKeysHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/keys/signed",
					Handler: accounts.GetKeysSignedHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/v1/keys/signed",
					Handler: accounts.PutKeysSignedHandler(serverCtx),
				},
			}...,
		),
	)

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckBasicAuth},
			[]rest.Route{
				{
					Method:  http.MethodPut,
					Path:    "/v1/messages/:destination",
					Handler: msgs.PutMsgsSendHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/messages",
					Handler: msgs.PostMsgsPendingHandler(serverCtx),
				},
			}...,
		),
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/textsecret/login",
				Handler: textsecret.AdxUserLoginHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserNameCheck},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/textsecret/ws",
					Handler: textsecret.AdxUserWSHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
