package protocol

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/app/middleware"
	"github.com/luyasr/simple-blog/config"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"net/http"
)

type HttpServer struct {
	server *http.Server
}

func NewHttpServer() *HttpServer {
	var ginLogMode string
	// 启动Gin, 注册路由
	if config.C.Server.Debug {
		ginLogMode = gin.DebugMode
	} else {
		ginLogMode = gin.ReleaseMode
	}
	gin.SetMode(ginLogMode)
	r := gin.New()
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	apiV1 := r.Group("api/v1")
	ioc.HttpHandler().RouteRegistry(apiV1)

	return &HttpServer{
		server: &http.Server{
			Addr:    config.C.Server.Addr(),
			Handler: r,
		},
	}
}

func (s *HttpServer) Run() error {
	_ = fmt.Sprintf("http server listen on %s", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
