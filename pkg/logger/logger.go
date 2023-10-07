package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

var (
	Console = NewConsoleLog()
)

func init() {
	zerolog.TimeFieldFormat = time.DateTime
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

func NewConsoleLog() zerolog.Logger {
	return zerolog.New(console()).With().Timestamp().Logger()
}

func NewFileLog(name string) zerolog.Logger {
	return zerolog.New(file(name)).With().Timestamp().Caller().Logger()
}

func NewMultiLog(name string) zerolog.Logger {
	multi := zerolog.MultiLevelWriter(console(), file(name))
	return zerolog.New(multi).With().Timestamp().Logger()
}

func console() zerolog.ConsoleWriter {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s |", i))
	}

	return output
}

func rootPath() string {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(path.Dir(filename)))
	return root
}

func file(name string) *os.File {
	var err error
	var dir string
	if config.C.Server.Log.Dir == "" {
		err = os.MkdirAll(fmt.Sprintf("%s/log/%s", rootPath(), name), os.ModePerm)
		dir = fmt.Sprintf("%s/log", rootPath())
	} else {
		err = os.MkdirAll(fmt.Sprintf("%s/%s", config.C.Server.Log.Dir, name), os.ModePerm)
		dir = config.C.Server.Log.Dir
	}
	if err != nil {
		panic(err)
	}

	filename := fmt.Sprintf("%s/%s/%s.log", dir, name, time.Now().Format("2006-01-02-15"))
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}

	return file
}

// GinLogger receive gin framework default log
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start)

		Console.Info().
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
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					Console.Error().
						Any("errors", err).
						Str("request", string(httpRequest)).Send()
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: err check
					c.Abort()
					return
				}

				if stack {
					Console.Error().
						Any("errors", err).
						Str("request", string(httpRequest)).
						Str("stack", string(debug.Stack())).Send()
				} else {
					Console.Error().
						Any("errors", err).
						Str("request", string(httpRequest)).Send()
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
