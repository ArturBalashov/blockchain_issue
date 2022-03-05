package blockchain

import (
	"crypto/sha256"
	"encoding"
	"hash"
	"math/rand"
)

const (
	SolutionSize = 1
	PuzzleSize   = 1
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
	return &PuzzleSolver{
		Puzzle:          p,
		PrecomputedHash: sha256.New(),
	}
}

func (ps *PuzzleSolver) Solve() SolvingResult {
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
	hasher := ps.PrecomputedHash
	hasher.Write(solution)
	hasher.Sum(nil)
	marshaller, ok := hasher.(encoding.BinaryMarshaler)
	if !ok {
		return false
	}
	state, err := marshaller.MarshalBinary()
	if err != nil {
		return false
	}

	leadingZeros := 0

	for i := 0; i < int(ps.Puzzle.Complexity)/2+1; i++ {
		c := state[i]
		if c>>4 == 0 {
			leadingZeros += 1
		} else {
			break
		}

		if c&0xF == 0 {
			leadingZeros += 1
		} else {
			break
		}
	}

	return leadingZeros >= int(ps.Puzzle.Complexity)
}
