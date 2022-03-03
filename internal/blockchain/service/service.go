package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/ArturBalashov/blockchain_issue/internal/blockchain/config"
	pb "github.com/ArturBalashov/blockchain_issue/internal/blockchain/pb"
	"github.com/ArturBalashov/blockchain_issue/internal/repository"
)

func New(repo repository.Repository, logger *zap.Logger, cfg *config.Config) *serviceWisdoms {
	return &serviceWisdoms{
		repo:   repo,
		logger: logger,
		cfg:    cfg,
	}
}

type serviceWisdoms struct {
	repo   repository.Repository
	logger *zap.Logger
	cfg    *config.Config
}

func (s *serviceWisdoms) GetIssue(_ context.Context, _ *pb.GetIssueRequest) (*pb.GetIssueResponse, error) {
	return nil, nil
}

func (s *serviceWisdoms) GetQuote(_ context.Context, _ *pb.GetQuoteRequest) (*pb.GetQuoteResponse, error) {
	return &pb.GetQuoteResponse{
		Quote: s.repo.GetQuote(),
	}, nil
}
