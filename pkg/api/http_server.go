package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/config"
	"github.com/luyasr/simple-blog/pkg/logger"
	"github.com/luyasr/simple-blog/pkg/token"
	"github.com/luyasr/simple-blog/pkg/user"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() {
	var ginLogMode string
	if config.C.Server.Debug {
		ginLogMode = gin.DebugMode
	} else {
		ginLogMode = gin.ReleaseMode
	}
	gin.SetMode(ginLogMode)
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	api := r.Group("api/v1")

	// 注册到主路由
	userServiceImpl := user.NewServiceImpl()
	u := user.NewHandler()

	t := token.NewHandler(userServiceImpl)
	api.Use()
	{
		InitRoute(api)
		u.InitUserRoute(api)
		t.InitTokenRoute(api)
	}

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
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
