// Package cmd /*
package cmd

import (
	"context"
	"errors"
	"github.com/luyasr/simple-blog/pkg/logger"
	"github.com/luyasr/simple-blog/protocol"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "simple-blog API服务",
	Long:  "simple-blog API服务",
	Run: func(cmd *cobra.Command, args []string) {
		server := newServer()
		server.start()
	},
}

type Server struct {
	http *protocol.HttpServer
	ch   chan os.Signal
	log  zerolog.Logger
}

func newServer() *Server {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	return &Server{
		http: protocol.NewHttpServer(),
		ch:   sigs,
		log:  logger.NewConsoleLog(),
	}
}

func (s *Server) start() {
	go func() {
		err := s.http.Run()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
	s.waitSigs(s.ch)
}

func (s *Server) waitSigs(sigs chan os.Signal) {
	<-sigs
	s.log.Info().Msg("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.http.Shutdown(ctx); err != nil {
		s.log.Error().Err(err).Msgf("Shutdown Server Error: ", err)
	}
	s.log.Info().Msg("Server exiting")
}
