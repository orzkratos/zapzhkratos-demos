# Changes

Code differences compared to source project demokratos.

## cmd/demo1kratos/main.go (+13 -14)

```diff
@@ -7,14 +7,16 @@
 	"github.com/go-kratos/kratos/v2"
 	"github.com/go-kratos/kratos/v2/config"
 	"github.com/go-kratos/kratos/v2/config/file"
-	"github.com/go-kratos/kratos/v2/log"
-	"github.com/go-kratos/kratos/v2/middleware/tracing"
 	"github.com/go-kratos/kratos/v2/transport/grpc"
 	"github.com/go-kratos/kratos/v2/transport/http"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/conf"
+	"github.com/orzkratos/zapzhkratos"
 	"github.com/yyle88/done"
 	"github.com/yyle88/must"
 	"github.com/yyle88/rese"
+	"github.com/yyle88/zaplog"
+	_ "go.uber.org/automaxprocs"
+	"go.uber.org/zap"
 )
 
 // go build -ldflags "-X main.Version=x.y.z"
@@ -31,13 +33,13 @@
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
+		kratos.Logger(zap匝普日志.New奎沱日志("网络服务")),
 		kratos.Server(
 			gs,
 			hs,
@@ -47,15 +49,12 @@
 
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
+	zapKratos := zapzhkratos.New匝普日志(zaplog.LOGGER, zapzhkratos.New日志配置())
+	zapLog := zapKratos.Sub模块匝普()
+	zapLog.LOG.Info("服务版本信息", zap.String("version", Version))
+	zapLog.LOG.Info("准备读取配置", zap.String("config", flagconf))
+
 	c := config.New(
 		config.WithSource(
 			file.NewSource(flagconf),
@@ -68,7 +67,7 @@
 	var cfg conf.Bootstrap
 	must.Done(c.Scan(&cfg))
 
-	app, cleanup := rese.V2(wireApp(cfg.Server, cfg.Data, logger))
+	app, cleanup := rese.V2(wireApp(cfg.Server, cfg.Data, zapKratos))
 	defer cleanup()
 
 	// start and wait for stop signal
```

## cmd/demo1kratos/wire.go (+2 -2)

```diff
@@ -6,16 +6,16 @@
 
 import (
 	"github.com/go-kratos/kratos/v2"
-	"github.com/go-kratos/kratos/v2/log"
 	"github.com/google/wire"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/biz"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/conf"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/data"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/server"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/service"
+	"github.com/orzkratos/zapzhkratos"
 )
 
 // wireApp init kratos application.
-func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
+func wireApp(*conf.Server, *conf.Data, *zapzhkratos.T匝普日志) (*kratos.App, func(), error) {
 	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
 }
```

## cmd/demo1kratos/wire_gen.go (+13 -9)

```diff
@@ -7,28 +7,32 @@
 
 import (
 	"github.com/go-kratos/kratos/v2"
-	"github.com/go-kratos/kratos/v2/log"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/biz"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/conf"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/data"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/server"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/service"
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

## internal/biz/greeter.go (+10 -6)

```diff
@@ -4,8 +4,9 @@
 	"context"
 
 	"github.com/go-kratos/kratos/v2/errors"
-	"github.com/go-kratos/kratos/v2/log"
 	v1 "github.com/orzkratos/demokratos/demo1kratos/api/helloworld/v1"
+	"github.com/orzkratos/zapzhkratos"
+	"github.com/yyle88/zaplog"
 )
 
 var (
@@ -29,17 +30,20 @@
 
 // GreeterUsecase is a Greeter usecase.
 type GreeterUsecase struct {
-	repo GreeterRepo
-	log  *log.Helper
+	repo   GreeterRepo
+	zapLog *zaplog.Zap
 }
 
 // NewGreeterUsecase new a Greeter usecase.
-func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
-	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
+func NewGreeterUsecase(repo GreeterRepo, zap匝普日志 *zapzhkratos.T匝普日志) *GreeterUsecase {
+	return &GreeterUsecase{
+		repo:   repo,
+		zapLog: zap匝普日志.Sub模块匝普(),
+	}
 }
 
 // CreateGreeter creates a Greeter, and returns the new Greeter.
 func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
-	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
+	uc.zapLog.SUG.Infof("存储打招呼信息: %v", g.Hello)
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
 	"github.com/orzkratos/demokratos/demo1kratos/internal/conf"
+	"github.com/orzkratos/zapzhkratos"
 )
 
 // ProviderSet is data providers.
@@ -15,9 +15,11 @@
 }
 
 // NewData .
-func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
+func NewData(c *conf.Data, zap匝普日志 *zapzhkratos.T匝普日志) (*Data, func(), error) {
+	zapLog := zap匝普日志.Sub模块匝普()
+	zapLog.SUG.Debugln("准备链接数据资源")
 	cleanup := func() {
-		log.NewHelper(logger).Info("closing the data resources")
+		zapLog.SUG.Info("准备关闭数据资源")
 	}
 	return &Data{}, cleanup, nil
 }
```

## internal/data/greeter.go (+9 -6)

```diff
@@ -3,24 +3,27 @@
 import (
 	"context"
 
-	"github.com/go-kratos/kratos/v2/log"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/biz"
+	"github.com/orzkratos/zapzhkratos"
+	"github.com/yyle88/zaplog"
+	"go.uber.org/zap"
 )
 
 type greeterRepo struct {
-	data *Data
-	log  *log.Helper
+	data   *Data
+	zapLog *zaplog.Zap
 }
 
 // NewGreeterRepo .
-func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
+func NewGreeterRepo(data *Data, zap匝普日志 *zapzhkratos.T匝普日志) biz.GreeterRepo {
 	return &greeterRepo{
-		data: data,
-		log:  log.NewHelper(logger),
+		data:   data,
+		zapLog: zap匝普日志.Sub模块匝普(),
 	}
 }
 
 func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
+	r.zapLog.LOG.Info("保存打招呼信息", zap.String("hello", g.Hello))
 	return g, nil
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
 	v1 "github.com/orzkratos/demokratos/demo1kratos/api/helloworld/v1"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/conf"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/service"
+	"github.com/orzkratos/zapzhkratos"
 )
 
 // NewGRPCServer new a gRPC server.
-func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *grpc.Server {
+func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, zap匝普日志 *zapzhkratos.T匝普日志) *grpc.Server {
 	var opts = []grpc.ServerOption{
 		grpc.Middleware(
 			recovery.Recovery(),
+			logging.Server(zap匝普日志.Get奎沱日志("GRPC请求日志")),
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
 	v1 "github.com/orzkratos/demokratos/demo1kratos/api/helloworld/v1"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/conf"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/service"
+	"github.com/orzkratos/zapzhkratos"
 )
 
 // NewHTTPServer new an HTTP server.
-func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
+func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, zap匝普日志 *zapzhkratos.T匝普日志) *http.Server {
 	var opts = []http.ServerOption{
 		http.Middleware(
 			recovery.Recovery(),
+			logging.Server(zap匝普日志.Get奎沱日志("HTTP请求日志")),
 		),
 	}
 	if c.Http.Network != "" {
```

## internal/service/greeter.go (+12 -3)

```diff
@@ -5,25 +5,34 @@
 
 	v1 "github.com/orzkratos/demokratos/demo1kratos/api/helloworld/v1"
 	"github.com/orzkratos/demokratos/demo1kratos/internal/biz"
+	"github.com/orzkratos/zapzhkratos"
+	"github.com/yyle88/zaplog"
+	"go.uber.org/zap"
 )
 
 // GreeterService is a greeter service.
 type GreeterService struct {
 	v1.UnimplementedGreeterServer
 
-	uc *biz.GreeterUsecase
+	uc     *biz.GreeterUsecase
+	zapLog *zaplog.Zap
 }
 
 // NewGreeterService new a greeter service.
-func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
-	return &GreeterService{uc: uc}
+func NewGreeterService(uc *biz.GreeterUsecase, zap匝普日志 *zapzhkratos.T匝普日志) *GreeterService {
+	return &GreeterService{
+		uc:     uc,
+		zapLog: zap匝普日志.Sub模块匝普(),
+	}
 }
 
 // SayHello implements helloworld.GreeterServer.
 func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
+	s.zapLog.LOG.Info("收到打招呼信息", zap.String("name", in.Name))
 	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
 	if err != nil {
 		return nil, err
 	}
+	s.zapLog.LOG.Info("回复打招呼信息", zap.String("name", in.Name))
 	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
 }
```

