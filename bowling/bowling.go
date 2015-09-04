package bowling

type game struct {
	frames []frame
	score  int
}

type frame struct {
	rolls []int
}

var rolls = make([]int, 22)
var rollIndex = 0

func (g *game) ScoreFrame(ballOne int, ballTwo int) {
	rolls[rollIndex] = ballOne
	rolls[rollIndex+1] = ballTwo
	g.frames = append(g.frames, frame{rolls[rollIndex : rollIndex+2]})
	rollIndex = rollIndex + 2
}

func (g game) GameScore() int {
	var score int
	for _, frame := range g.frames {
		score = score + frame.FrameScore()
	}
	return score
}

func (f frame) FrameScore() int {
	return f.rolls[0] + f.rolls[1]
}
