package highscores

type HighScores struct{
    top []int
    latest int
    scores []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
    s := HighScores{}
    s.scores = scores
    s.top = make([]int, 3)
    for _,score := range scores {
        s.addScore(score)
    }
    return &s
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.scores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.latest
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	return s.top[0]
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
    switch {
        case s.top[1] == 0:
        return s.top[:1]
        case s.top[2] == 0:
        return s.top[:2]
        default:
        return s.top
    }
}

func (s *HighScores) addScore(score int) {
    switch {
        case score > s.top[0]:
        	s.top[2] = s.top[1]
        	s.top[1] = s.top[0]
        	s.top[0] = score
        case score > s.top[1]:
        	s.top[2] = s.top[1]
        	s.top[1] = score
        case score > s.top[2]:
        	s.top[2] = score
    }
    s.latest = score
}