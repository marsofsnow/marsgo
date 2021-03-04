package config

import (
	"flag"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

var configFile = flag.String("f", "etc/api-api.yaml", "the config file")
var AppConfig Config


func init()  {
	flag.Parse()
	conf.MustLoad(*configFile, &AppConfig)
	logx.MustSetup(AppConfig.LogConf.LogConf)   //日志配置

}

type Config struct {
	LogConf struct{
		logx.LogConf
	}
	rest.RestConf
	Auth struct {
		AccessSecret string  // AccessSecret：生成jwt token的密钥，最简单的方式可以使用一个uuid值
		AccessExpire int64   // jwt token有效期，单位：秒
	}
	WssAddress string
	CacheRedis struct{
		Addr string
		Password string
		DB int
	}
	DirectoryRedis struct{
		Addr string
		Password string
		DB int
	}
	Mysql struct{
		DataSource string
	}
	BookStore zrpc.RpcClientConf

	TestDevices []struct{
		Number string
		Code string
	}
	EOSChainUrls []string

}
