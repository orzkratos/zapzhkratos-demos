//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/orzkratos/demokratos/demo1kratos/internal/biz"
	"github.com/orzkratos/demokratos/demo1kratos/internal/conf"
	"github.com/orzkratos/demokratos/demo1kratos/internal/data"
	"github.com/orzkratos/demokratos/demo1kratos/internal/server"
	"github.com/orzkratos/demokratos/demo1kratos/internal/service"
	"github.com/orzkratos/zapzhkratos"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *zapzhkratos.T匝普日志) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
