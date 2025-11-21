package main

import (
	"flag"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/orzkratos/demokratos/demo2kratos/internal/conf"
	"github.com/orzkratos/zapzhkratos"
	"github.com/yyle88/done"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
	_ "go.uber.org/automaxprocs"
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
		kratos.Logger(zap匝普日志.Get奎沱日志("网络服务")),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()

	// demo2 uses Get奎沱秘书 to get *log.Helper (Kratos style)
	// demo2 使用 Get奎沱秘书 获取 *log.Helper（Kratos 风格）
	zapKratos := zapzhkratos.New匝普日志(zaplog.LOGGER, zapzhkratos.New日志配置())
	slog := zapKratos.Get奎沱秘书("启动日志")
	slog.Infof("服务版本: %s", Version)
	slog.Infof("配置路径: %s", flagconf)

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer rese.F0(c.Close)

	must.Done(c.Load())

	var cfg conf.Bootstrap
	must.Done(c.Scan(&cfg))

	app, cleanup := rese.V2(wireApp(cfg.Server, cfg.Data, zapKratos))
	defer cleanup()

	// start and wait for stop signal
	must.Done(app.Run())
}
