package inmemory

import (
	"bufio"
	"math/rand"
	"os"
	"sync"

	"github.com/ArturBalashov/blockchain_issue/pkg/tools/blockchain"
	"github.com/ArturBalashov/blockchain_issue/pkg/tools/helpers"
)

var DefaultIdSize = 5
var mu sync.Mutex
var hashes = make(map[string]blockchain.Puzzle)

func New(filename string) (*db, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var quotes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		quotes = append(quotes, scanner.Text())
	}

	return &db{quotes: quotes}, nil
}

type db struct {
	quotes []string
}

func (d *db) GetQuote() string {
	num := rand.Intn(len(d.quotes))
	return d.quotes[num]
}

func (d *db) AddHash(puz *blockchain.Puzzle) string {
	uid := helpers.GenerateRandomString(DefaultIdSize)
	mu.Lock()
	hashes[uid] = *puz
	mu.Unlock()

	return uid
}

func (d *db) Deletehash(uid string) {
	mu.Lock()
	delete(hashes, uid)
	mu.Unlock()
}

func (d *db) GetHash(uid string) blockchain.Puzzle {
	return hashes[uid]
}
