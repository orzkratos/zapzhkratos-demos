package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	v1 "github.com/orzkratos/demokratos/demo1kratos/api/helloworld/v1"
	"github.com/orzkratos/zapzhkratos"
	"github.com/yyle88/zaplog"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Update(context.Context, *Greeter) (*Greeter, error)
	FindByID(context.Context, int64) (*Greeter, error)
	ListByHello(context.Context, string) ([]*Greeter, error)
	ListAll(context.Context) ([]*Greeter, error)
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo   GreeterRepo
	zapLog *zaplog.Zap
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, zap匝普日志 *zapzhkratos.T匝普日志) *GreeterUsecase {
	return &GreeterUsecase{
		repo:   repo,
		zapLog: zap匝普日志.Sub模块匝普(),
	}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.zapLog.SUG.Infof("存储打招呼信息: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
