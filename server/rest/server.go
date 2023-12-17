package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alvinfebriando/gin-gorm-skeleton/applogger"
	"github.com/alvinfebriando/gin-gorm-skeleton/config"
)

func New(router http.Handler) *http.Server {
	addr := fmt.Sprintf("%s:%s", config.New().AppHost, config.New().AppPort)

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

func StartWithGracefulShutdown(s *http.Server) {
	applogger.Log.Info("Starting server ...")
	go func() {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			applogger.Log.Fatalf("listen: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	applogger.Log.Info("Shutdown Server ...")

	timeout := config.New().GracefulShutdownTimeout

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		applogger.Log.Info("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		applogger.Log.Infof("timeout of %d seconds.", timeout)
	}

	applogger.Log.Info("server exiting")
}
