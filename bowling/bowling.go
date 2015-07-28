package bowling

type game struct {
	frames []frame
	score  int
}

type frame struct {
	firstBall  int
	secondBall int
}

func (g *game) ScoreFrame(ballOne int, ballTwo int) {
	g.frames = append(g.frames, frame{ballOne, ballTwo})
}

func (g game) GameScore() int {
	var score int
	for _, frame := range g.frames {
		score = score + frame.firstBall + frame.secondBall
	}
	return score
}
