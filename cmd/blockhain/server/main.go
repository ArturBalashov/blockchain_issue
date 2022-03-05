package main

import (
	"net"
	"os"
	"os/signal"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"

	"github.com/ArturBalashov/blockchain_issue/internal/blockchain/config"
	pb "github.com/ArturBalashov/blockchain_issue/internal/blockchain/pb"
	"github.com/ArturBalashov/blockchain_issue/internal/blockchain/service"
	"github.com/ArturBalashov/blockchain_issue/internal/repository"
)

func main() {
	// logger initialize
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stdout,
		zap.DebugLevel,
	))
	logger.Info("service starting...")

	defer logger.Info("service stopped")

	// config initialize
	cfg := config.New()
	repo, err := repository.New(cfg.FileResponsesPath)
	if err != nil {
		logger.Error("can't create repository", zap.Error(err))
		return
	}
	srv := service.New(repo, logger, cfg)

	grpcSrv := grpc.NewServer()
	pb.RegisterBlockchainServer(grpcSrv, srv)

	listener, err := net.Listen("tcp", cfg.SrvAddr)
	if err != nil {
		logger.Error("can't create listener", zap.Error(err))
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		sig := <-quit
		logger.Info("shutdown", zap.String("signal", sig.String()))
		grpcSrv.GracefulStop()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		logger.Info("service was starting, address http listening", zap.String("address", cfg.SrvAddr))
		if err := grpcSrv.Serve(listener); err != nil {
			logger.Error("http listen error", zap.Error(err))
			return
		}
	}()

	wg.Wait()
}
