package config

import (
	"github.com/ArturBalashov/blockchain_issue/pkg/tools/helpers"
)

func New() *Config {
	cfg := Config{}

	helpers.EnvToString(&cfg.SrvAddr, "BLOCKCHAIN_ADDRESS", ":8080")
	helpers.EnvToString(&cfg.FileResponsesPath, "BLOCKCHAIN_FILEPATH", "internal/repository/in_memory/quotes.txt")

	return &cfg
}

type Config struct {
	SrvAddr string

	FileResponsesPath string
}
