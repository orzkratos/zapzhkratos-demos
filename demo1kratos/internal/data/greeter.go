package data

import (
	"context"

	"github.com/orzkratos/demokratos/demo1kratos/internal/biz"
	"github.com/orzkratos/zapzhkratos"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type greeterRepo struct {
	data   *Data
	zapLog *zaplog.Zap
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, zap匝普日志 *zapzhkratos.T匝普日志) biz.GreeterRepo {
	return &greeterRepo{
		data:   data,
		zapLog: zap匝普日志.Sub模块匝普(),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	r.zapLog.LOG.Info("保存打招呼信息", zap.String("hello", g.Hello))
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}
