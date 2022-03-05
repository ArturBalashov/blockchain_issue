package repository

import (
	inmemory "github.com/ArturBalashov/blockchain_issue/internal/repository/in_memory"
	"github.com/ArturBalashov/blockchain_issue/pkg/tools/blockchain"
)

func New(filename string) (Repository, error) {
	return inmemory.New(filename)
}

type Repository interface {
	GetQuote() string
	AddHash(puz *blockchain.Puzzle) string
	Deletehash(uid string)
	GetHash(uid string) blockchain.Puzzle
}
