package repository

import inmemory "github.com/ArturBalashov/blockchain_issue/internal/repository/in_memory"

func New(filename string) (Repository, error) {
	return inmemory.New(filename)
}

type Repository interface {
	GetQuote() string
}
