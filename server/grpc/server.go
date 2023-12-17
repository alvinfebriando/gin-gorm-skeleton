package grpc

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alvinfebriando/gin-gorm-skeleton/applogger"
	"github.com/alvinfebriando/gin-gorm-skeleton/config"
	grpchandler "github.com/alvinfebriando/gin-gorm-skeleton/handler/grpc"
	"github.com/alvinfebriando/gin-gorm-skeleton/interceptor"
	pb "github.com/alvinfebriando/gin-gorm-skeleton/proto"
	"google.golang.org/grpc"
)

type Handlers struct {
	User *grpchandler.UserHandler
}

func New(handlers Handlers) *grpc.Server {
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(
		interceptor.LogInterceptor,
		interceptor.AuthInterceptor,
	))

	pb.RegisterUserServer(s, handlers.User)

	return s
}

func StartWithGracefulShutdown(s *grpc.Server) {
	cfg := config.New()
	address := fmt.Sprintf("%s:%s", cfg.GrpcHost, cfg.GrpcPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		applogger.Log.Errorf("failed to listen: %v", err)
	}

	applogger.Log.Info("Starting server ...")
	go func() {
		if err = s.Serve(listener); err != nil {
			applogger.Log.Errorf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	applogger.Log.Info("Shutdown Server ...")

	stopped := make(chan struct{})
	go func() {
		s.GracefulStop()
		close(stopped)
	}()

	t := time.NewTimer(10 * time.Second)
	select {
	case <-t.C:
		s.Stop()
	case <-stopped:
		t.Stop()
	}

	applogger.Log.Info("server exiting")
}
