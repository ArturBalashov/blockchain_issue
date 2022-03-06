package blockchain

import (
	"crypto/sha256"
	"hash"
	"math/rand"
	"time"
)

const (
	SolutionSize = 16
	PuzzleSize   = 16
)

type PuzzleSolution []byte

type SolvingResult struct {
	Solution  PuzzleSolution
	HashTried int
}

type Puzzle struct {
	Complexity uint8
	Value      []byte
}

func NewPuzzle(complexity uint8) *Puzzle {
	value := make([]byte, PuzzleSize)
	rand.Read(value)
	return &Puzzle{
		Complexity: complexity,
		Value:      value,
	}
}

type PuzzleSolver struct {
	Puzzle          *Puzzle
	PrecomputedHash hash.Hash
}

func NewPuzzleSolver(p *Puzzle) *PuzzleSolver {
	preHash := sha256.New()
	preHash.Write(p.Value)
	preHash.Sum(nil)
	return &PuzzleSolver{
		Puzzle:          p,
		PrecomputedHash: preHash,
	}
}

func (ps *PuzzleSolver) Solve() SolvingResult {
	rand.Seed(time.Now().UnixNano())
	hashesTries := 0

	for {
		solution := make([]byte, SolutionSize)
		rand.Read(solution)
		hashesTries++
		if ps.IsValidSolution(solution) {
			return SolvingResult{
				Solution:  PuzzleSolution(solution),
				HashTried: hashesTries,
			}
		}
	}
}

func (ps *PuzzleSolver) IsValidSolution(solution PuzzleSolution) bool {
	hash := sha256.Sum256(solution)

	leadingZeros := 0

	for i := 0; i < (int(ps.Puzzle.Complexity)/2 + 1); i++ {
		if hash[i]>>4 == 0 {
			leadingZeros += 1
		} else {
			break
		}

		if hash[i]&0xF == 0 {
			leadingZeros += 1
		} else {
			break
		}
	}

	return leadingZeros >= int(ps.Puzzle.Complexity)
}
