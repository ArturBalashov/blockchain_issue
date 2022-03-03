package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/ArturBalashov/blockchain_issue/internal/blockchain/config"
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

	fmt.Println(srv)
}
