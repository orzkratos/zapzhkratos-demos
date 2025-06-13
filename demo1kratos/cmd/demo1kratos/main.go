package main

import (
	"flag"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/orzkratos/demokratos/demo1kratos/internal/conf"
	"github.com/orzkratos/zapzhkratos"
	"github.com/yyle88/done"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "./configs", "config path, eg: -conf config.yaml")
}

func newApp(gs *grpc.Server, hs *http.Server, zap匝普日志 *zapzhkratos.T匝普日志) *kratos.App {
	return kratos.New(
		kratos.ID(done.VCE(os.Hostname()).Omit()),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(zap匝普日志.New奎沱日志("网络服务")),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()

	zapKratos := zapzhkratos.New匝普日志(zaplog.LOGGER, zapzhkratos.New日志配置())
	zapLog := zapKratos.Sub模块匝普()
	zapLog.LOG.Info("服务版本信息", zap.String("version", Version))
	zapLog.LOG.Info("准备读取配置", zap.String("config", flagconf))

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer rese.F0(c.Close)

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, zapKratos)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
