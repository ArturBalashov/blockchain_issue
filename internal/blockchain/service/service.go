package service

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/ArturBalashov/blockchain_issue/internal/blockchain/config"
	pb "github.com/ArturBalashov/blockchain_issue/internal/blockchain/pb"
	"github.com/ArturBalashov/blockchain_issue/internal/repository"
	puzzle "github.com/ArturBalashov/blockchain_issue/pkg/tools/blockchain"
)

const DefaultComplexity = 1

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

func (s *serviceWisdoms) GetIssue(_ context.Context, req *pb.GetIssueRequest) (*pb.GetIssueResponse, error) {
	puz := puzzle.NewPuzzle(uint8(DefaultComplexity))
	uid := s.repo.AddHash(puz)
	return &pb.GetIssueResponse{
		Puzzle:     puz.Value,
		Complexity: int32(puz.Complexity),
		Uid:        uid,
	}, nil
}

func (s *serviceWisdoms) GetQuote(_ context.Context, req *pb.GetQuoteRequest) (*pb.GetQuoteResponse, error) {
	solution := puzzle.PuzzleSolution(req.Solution)
	puz := s.repo.GetHash(req.Uid)

	solver := puzzle.NewPuzzleSolver(&puz)
	if !solver.IsValidSolution(solution) {
		return nil, errors.New("worng solution")
	}

	return &pb.GetQuoteResponse{
		Quote: s.repo.GetQuote(),
	}, nil
}
