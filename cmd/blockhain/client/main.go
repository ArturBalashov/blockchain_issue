package main

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ArturBalashov/blockchain_issue/internal/blockchain/config"
	pb "github.com/ArturBalashov/blockchain_issue/internal/blockchain/pb"
	puzzle "github.com/ArturBalashov/blockchain_issue/pkg/tools/blockchain"
)

func main() {
	// logger initialize
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stdout,
		zap.DebugLevel,
	))
	logger.Info("client starting...")

	defer logger.Info("service stopped")

	// config initialize
	cfg := config.New()
	conn, err := grpc.Dial(cfg.SrvAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("can't conn to server", zap.Error(err))
		return
	}

	client := pb.NewBlockchainClient(conn)

	issue, err := client.GetIssue(context.Background(), &pb.GetIssueRequest{})
	if err != nil {
		logger.Error("can't get an issue", zap.Error(err))
		return
	}

	puz := puzzle.Puzzle{Value: issue.GetPuzzle(), Complexity: uint8(issue.GetComplexity())}
	logger.Info("puzzle", zap.Binary("value", puz.Value), zap.Uint8("complexity", puz.Complexity))
	solver := puzzle.NewPuzzleSolver(&puz)

	logger.Info("solving...")
	solution := solver.Solve()
	logger.Info("solution", zap.Binary("solution", solution.Solution), zap.Int("hash_tried", solution.HashTried))

	quote, err := client.GetQuote(context.Background(), &pb.GetQuoteRequest{Solution: solution.Solution})
	if err != nil {
		logger.Error("can't get a quote", zap.Error(err))
		return
	}

	logger.Info("response", zap.String("quote", quote.GetQuote()))
}
