package highscores

import (
	"cmp"
	"math"
	"slices"
)

type HighScores struct {
	scores []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	return &HighScores{
		scores: scores,
	}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.scores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	if len(s.scores) == 0 {
		return 0
	}
	return s.scores[len(s.scores)-1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	if len(s.scores) == 0 {
		return 0 // Or handle error/panic
	}
	
	maxScore := math.MinInt
	for i := 0; i < len(s.scores); i++ {
		maxScore = max(maxScore, s.scores[i])
	}
	return maxScore
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	sortedSlice := make([]int, len(s.scores))
	copy(sortedSlice, s.scores)

	slices.SortFunc(sortedSlice, func(a, b int) int {
		return cmp.Compare(b, a) // Descending order
	})

	if len(sortedSlice) < 3 {
		return sortedSlice
	}
	return sortedSlice[0:3]
}   