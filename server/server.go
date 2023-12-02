package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/alvinfebriando/gin-gorm-skeleton/applogger"
)

func New(router http.Handler) *http.Server {
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	addr := fmt.Sprintf("%s:%s", host, port)

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

func StartWithGracefulShutdown(s *http.Server) {
	go func() {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			applogger.Log.Fatalf("listen: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	applogger.Log.Info("Shutdown Server ...")

	const defaultTimeout = 5
	timeoutString := os.Getenv("TIMEOUT")
	timeout, err := strconv.Atoi(timeoutString)
	if err != nil {
		timeout = defaultTimeout
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	if err = s.Shutdown(ctx); err != nil {
		applogger.Log.Info("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		applogger.Log.Infof("timeout of %d seconds.", timeout)
	}

	applogger.Log.Info("server exiting")
}
