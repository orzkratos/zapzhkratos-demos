package data

import (
	"github.com/google/wire"
	"github.com/orzkratos/demokratos/demo1kratos/internal/conf"
	"github.com/orzkratos/zapzhkratos"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, zap匝普日志 *zapzhkratos.T匝普日志) (*Data, func(), error) {
	zapLog := zap匝普日志.Sub模块匝普()
	zapLog.SUG.Debugln("准备链接数据资源")
	cleanup := func() {
		zapLog.SUG.Info("准备关闭数据资源")
	}
	return &Data{}, cleanup, nil
}
