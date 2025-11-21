package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/orzkratos/demokratos/demo2kratos/internal/biz"
	"github.com/orzkratos/zapzhkratos"
)

type greeterRepo struct {
	data *Data
	slog *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, zap匝普日志 *zapzhkratos.T匝普日志) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		slog: zap匝普日志.Get奎沱秘书("数据仓库"),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
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
