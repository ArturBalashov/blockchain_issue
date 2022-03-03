package inmemory

import (
	"bufio"
	"math/rand"
	"os"
)

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
