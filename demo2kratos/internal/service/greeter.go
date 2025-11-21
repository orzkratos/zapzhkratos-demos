package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/orzkratos/demokratos/demo2kratos/api/helloworld/v1"
	"github.com/orzkratos/demokratos/demo2kratos/internal/biz"
	"github.com/orzkratos/zapzhkratos"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc   *biz.GreeterUsecase
	slog *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, zap匝普日志 *zapzhkratos.T匝普日志) *GreeterService {
	return &GreeterService{
		uc:   uc,
		slog: zap匝普日志.Get奎沱秘书("服务层"),
	}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.slog.WithContext(ctx).Infof("收到请求: %s", in.Name)
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	s.slog.WithContext(ctx).Infof("返回响应: %s", g.Hello)
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
