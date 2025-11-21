# Changes

Code differences compared to source project demokratos.

## cmd/demo2kratos/main.go (+14 -14)

```diff
@@ -7,14 +7,15 @@
 	"github.com/go-kratos/kratos/v2"
 	"github.com/go-kratos/kratos/v2/config"
 	"github.com/go-kratos/kratos/v2/config/file"
-	"github.com/go-kratos/kratos/v2/log"
-	"github.com/go-kratos/kratos/v2/middleware/tracing"
 	"github.com/go-kratos/kratos/v2/transport/grpc"
 	"github.com/go-kratos/kratos/v2/transport/http"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/conf"
+	"github.com/orzkratos/zapzhkratos"
 	"github.com/yyle88/done"
 	"github.com/yyle88/must"
 	"github.com/yyle88/rese"
+	"github.com/yyle88/zaplog"
+	_ "go.uber.org/automaxprocs"
 )
 
 // go build -ldflags "-X main.Version=x.y.z"
@@ -31,13 +32,13 @@
 	flag.StringVar(&flagconf, "conf", "./configs", "config path, eg: -conf config.yaml")
 }
 
-func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
+func newApp(gs *grpc.Server, hs *http.Server, zap匝普日志 *zapzhkratos.T匝普日志) *kratos.App {
 	return kratos.New(
 		kratos.ID(done.VCE(os.Hostname()).Omit()),
 		kratos.Name(Name),
 		kratos.Version(Version),
 		kratos.Metadata(map[string]string{}),
-		kratos.Logger(logger),
+		kratos.Logger(zap匝普日志.Get奎沱日志("网络服务")),
 		kratos.Server(
 			gs,
 			hs,
@@ -47,15 +48,14 @@
 
 func main() {
 	flag.Parse()
-	logger := log.With(log.NewStdLogger(os.Stdout),
-		"ts", log.DefaultTimestamp,
-		"caller", log.DefaultCaller,
-		"service.id", kratos.ID(done.VCE(os.Hostname()).Omit()),
-		"service.name", Name,
-		"service.version", Version,
-		"trace.id", tracing.TraceID(),
-		"span.id", tracing.SpanID(),
-	)
+
+	// demo2 uses Get奎沱秘书 to get *log.Helper (Kratos style)
+	// demo2 使用 Get奎沱秘书 获取 *log.Helper（Kratos 风格）
+	zapKratos := zapzhkratos.New匝普日志(zaplog.LOGGER, zapzhkratos.New日志配置())
+	slog := zapKratos.Get奎沱秘书("启动日志")
+	slog.Infof("服务版本: %s", Version)
+	slog.Infof("配置路径: %s", flagconf)
+
 	c := config.New(
 		config.WithSource(
 			file.NewSource(flagconf),
@@ -68,7 +68,7 @@
 	var cfg conf.Bootstrap
 	must.Done(c.Scan(&cfg))
 
-	app, cleanup := rese.V2(wireApp(cfg.Server, cfg.Data, logger))
+	app, cleanup := rese.V2(wireApp(cfg.Server, cfg.Data, zapKratos))
 	defer cleanup()
 
 	// start and wait for stop signal
```

## cmd/demo2kratos/wire.go (+2 -2)

```diff
@@ -6,16 +6,16 @@
 
 import (
 	"github.com/go-kratos/kratos/v2"
-	"github.com/go-kratos/kratos/v2/log"
 	"github.com/google/wire"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/biz"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/conf"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/data"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/server"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/service"
+	"github.com/orzkratos/zapzhkratos"
 )
 
 // wireApp init kratos application.
-func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
+func wireApp(*conf.Server, *conf.Data, *zapzhkratos.T匝普日志) (*kratos.App, func(), error) {
 	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
 }
```

## cmd/demo2kratos/wire_gen.go (+13 -9)

```diff
@@ -7,28 +7,32 @@
 
 import (
 	"github.com/go-kratos/kratos/v2"
-	"github.com/go-kratos/kratos/v2/log"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/biz"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/conf"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/data"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/server"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/service"
+	"github.com/orzkratos/zapzhkratos"
 )
 
+import (
+	_ "go.uber.org/automaxprocs"
+)
+
 // Injectors from wire.go:
 
 // wireApp init kratos application.
-func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
-	dataData, cleanup, err := data.NewData(confData, logger)
+func wireApp(confServer *conf.Server, confData *conf.Data, t匝普日志 *zapzhkratos.T匝普日志) (*kratos.App, func(), error) {
+	dataData, cleanup, err := data.NewData(confData, t匝普日志)
 	if err != nil {
 		return nil, nil, err
 	}
-	greeterRepo := data.NewGreeterRepo(dataData, logger)
-	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
-	greeterService := service.NewGreeterService(greeterUsecase)
-	grpcServer := server.NewGRPCServer(confServer, greeterService, logger)
-	httpServer := server.NewHTTPServer(confServer, greeterService, logger)
-	app := newApp(logger, grpcServer, httpServer)
+	greeterRepo := data.NewGreeterRepo(dataData, t匝普日志)
+	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, t匝普日志)
+	greeterService := service.NewGreeterService(greeterUsecase, t匝普日志)
+	grpcServer := server.NewGRPCServer(confServer, greeterService, t匝普日志)
+	httpServer := server.NewHTTPServer(confServer, greeterService, t匝普日志)
+	app := newApp(grpcServer, httpServer, t匝普日志)
 	return app, func() {
 		cleanup()
 	}, nil
```

## internal/biz/greeter.go (+5 -4)

```diff
@@ -6,6 +6,7 @@
 	"github.com/go-kratos/kratos/v2/errors"
 	"github.com/go-kratos/kratos/v2/log"
 	v1 "github.com/orzkratos/demokratos/demo2kratos/api/helloworld/v1"
+	"github.com/orzkratos/zapzhkratos"
 )
 
 var (
@@ -30,16 +31,16 @@
 // GreeterUsecase is a Greeter usecase.
 type GreeterUsecase struct {
 	repo GreeterRepo
-	log  *log.Helper
+	slog *log.Helper
 }
 
 // NewGreeterUsecase new a Greeter usecase.
-func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
-	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
+func NewGreeterUsecase(repo GreeterRepo, zap匝普日志 *zapzhkratos.T匝普日志) *GreeterUsecase {
+	return &GreeterUsecase{repo: repo, slog: zap匝普日志.Get奎沱秘书("业务逻辑")}
 }
 
 // CreateGreeter creates a Greeter, and returns the new Greeter.
 func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
-	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
+	uc.slog.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
 	return uc.repo.Save(ctx, g)
 }
```

## internal/data/data.go (+5 -3)

```diff
@@ -1,9 +1,9 @@
 package data
 
 import (
-	"github.com/go-kratos/kratos/v2/log"
 	"github.com/google/wire"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/conf"
+	"github.com/orzkratos/zapzhkratos"
 )
 
 // ProviderSet is data providers.
@@ -15,9 +15,11 @@
 }
 
 // NewData .
-func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
+func NewData(c *conf.Data, zap匝普日志 *zapzhkratos.T匝普日志) (*Data, func(), error) {
+	slog := zap匝普日志.Get奎沱秘书("数据层")
+	slog.Debug("准备链接数据资源")
 	cleanup := func() {
-		log.NewHelper(logger).Info("closing the data resources")
+		slog.Info("准备关闭数据资源")
 	}
 	return &Data{}, cleanup, nil
 }
```

## internal/data/greeter.go (+4 -3)

```diff
@@ -5,18 +5,19 @@
 
 	"github.com/go-kratos/kratos/v2/log"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/biz"
+	"github.com/orzkratos/zapzhkratos"
 )
 
 type greeterRepo struct {
 	data *Data
-	log  *log.Helper
+	slog *log.Helper
 }
 
 // NewGreeterRepo .
-func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
+func NewGreeterRepo(data *Data, zap匝普日志 *zapzhkratos.T匝普日志) biz.GreeterRepo {
 	return &greeterRepo{
 		data: data,
-		log:  log.NewHelper(logger),
+		slog: zap匝普日志.Get奎沱秘书("数据仓库"),
 	}
 }
 
```

## internal/server/grpc.go (+4 -2)

```diff
@@ -1,19 +1,21 @@
 package server
 
 import (
-	"github.com/go-kratos/kratos/v2/log"
+	"github.com/go-kratos/kratos/v2/middleware/logging"
 	"github.com/go-kratos/kratos/v2/middleware/recovery"
 	"github.com/go-kratos/kratos/v2/transport/grpc"
 	v1 "github.com/orzkratos/demokratos/demo2kratos/api/helloworld/v1"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/conf"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/service"
+	"github.com/orzkratos/zapzhkratos"
 )
 
 // NewGRPCServer new a gRPC server.
-func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *grpc.Server {
+func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, zap匝普日志 *zapzhkratos.T匝普日志) *grpc.Server {
 	var opts = []grpc.ServerOption{
 		grpc.Middleware(
 			recovery.Recovery(),
+			logging.Server(zap匝普日志.Get奎沱日志("GRPC请求")),
 		),
 	}
 	if c.Grpc.Network != "" {
```

## internal/server/http.go (+4 -2)

```diff
@@ -1,19 +1,21 @@
 package server
 
 import (
-	"github.com/go-kratos/kratos/v2/log"
+	"github.com/go-kratos/kratos/v2/middleware/logging"
 	"github.com/go-kratos/kratos/v2/middleware/recovery"
 	"github.com/go-kratos/kratos/v2/transport/http"
 	v1 "github.com/orzkratos/demokratos/demo2kratos/api/helloworld/v1"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/conf"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/service"
+	"github.com/orzkratos/zapzhkratos"
 )
 
 // NewHTTPServer new an HTTP server.
-func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
+func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, zap匝普日志 *zapzhkratos.T匝普日志) *http.Server {
 	var opts = []http.ServerOption{
 		http.Middleware(
 			recovery.Recovery(),
+			logging.Server(zap匝普日志.Get奎沱日志("HTTP请求")),
 		),
 	}
 	if c.Http.Network != "" {
```

## internal/service/greeter.go (+11 -3)

```diff
@@ -3,27 +3,35 @@
 import (
 	"context"
 
+	"github.com/go-kratos/kratos/v2/log"
 	v1 "github.com/orzkratos/demokratos/demo2kratos/api/helloworld/v1"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/biz"
+	"github.com/orzkratos/zapzhkratos"
 )
 
 // GreeterService is a greeter service.
 type GreeterService struct {
 	v1.UnimplementedGreeterServer
 
-	uc *biz.GreeterUsecase
+	uc   *biz.GreeterUsecase
+	slog *log.Helper
 }
 
 // NewGreeterService new a greeter service.
-func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
-	return &GreeterService{uc: uc}
+func NewGreeterService(uc *biz.GreeterUsecase, zap匝普日志 *zapzhkratos.T匝普日志) *GreeterService {
+	return &GreeterService{
+		uc:   uc,
+		slog: zap匝普日志.Get奎沱秘书("服务层"),
+	}
 }
 
 // SayHello implements helloworld.GreeterServer.
 func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
+	s.slog.WithContext(ctx).Infof("收到请求: %s", in.Name)
 	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
 	if err != nil {
 		return nil, err
 	}
+	s.slog.WithContext(ctx).Infof("返回响应: %s", g.Hello)
 	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
 }
```

