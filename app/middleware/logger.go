package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/pkg/logger"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// GinLogger receive gin framework default log
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start)

		logger.Console.Info().
			Int("status", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("query", c.Request.URL.RawQuery).
			Str("ip", c.ClientIP()).
			Str("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()).
			Dur("cost", cost).
			Send()
	}
}

// GinRecovery Recover any panic that may occur in the project and use zero log to record relevant logs
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				brokenPipe := isBrokenPipe(err)

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Console.Error().
						Any("errors", err).
						Str("request", string(httpRequest)).Send()
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: err check
					c.Abort()
					return
				}

				if stack {
					logger.Console.Error().
						Any("errors", err).
						Str("request", string(httpRequest)).
						Str("stack", string(debug.Stack())).Send()
				} else {
					logger.Console.Error().
						Any("errors", err).
						Str("request", string(httpRequest)).Send()
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func isBrokenPipe(err any) bool {
	if ne, ok := err.(*net.OpError); ok {
		if se, ok := ne.Err.(*os.SyscallError); ok {
			if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
				return true
			}
		}
	}
	return false
}
