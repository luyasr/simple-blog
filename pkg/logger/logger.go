package logger

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
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
	return zerolog.New(rotate(name)).With().Timestamp().Caller().Logger()
}

func NewMultiLog(name string) zerolog.Logger {
	multi := zerolog.MultiLevelWriter(console(), rotate(name))
	return zerolog.New(multi).With().Timestamp().Logger()
}

func console() zerolog.ConsoleWriter {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s |", i))
	}

	return output
}

func rotate(name string) *lumberjack.Logger {
	logRotate := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s.log", name),
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}

	return logRotate
}
