package bowling

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScoreAllOnes(t *testing.T) {
	g := game{frames: make([]frame, 0)}
	g.ScoreFrame(1, 1)
	assert.Equal(t, 2, g.GameScore())
}

func TestTwoFrames(t *testing.T) {
	g := game{frames: make([]frame, 0)}
	g.ScoreFrame(2, 3)
	g.ScoreFrame(2, 5)
	assert.Equal(t, 12, g.GameScore())
}

/*
func TestSpare(t *testing.T) {
	g := game{frames: make([]frame, 10)}
	g.ScoreFrame(5, 5)
	g.ScoreFrame(8, 0)
	assert.Equal(t, 26, g.GameScore())
}
*/
