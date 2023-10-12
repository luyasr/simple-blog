package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/config"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/luyasr/simple-blog/pkg/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	var ginLogMode string

	// 初始化Controller
	if err := ioc.Controller().Init(); err != nil {
		panic(err)
	}

	// 初始化Handler
	if err := ioc.Handler().Init(); err != nil {
		panic(err)
	}

	// 启动Gin, 注册路由
	if config.C.Server.Debug {
		ginLogMode = gin.DebugMode
	} else {
		ginLogMode = gin.ReleaseMode
	}
	gin.SetMode(ginLogMode)
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	apiV1 := r.Group("api/v1")
	ioc.Handler().RouteRegistry(apiV1)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.C.Server.Port),
		Handler: r,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sigs
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
