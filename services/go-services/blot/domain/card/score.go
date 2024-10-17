package card

type Score struct {
	value int
}

func (s Score) Add(score Score) Score {
	return NewScore(s.value + score.value)
}

func (s Score) Value() int {
	return s.value
}

func (s Score) ConvertToTeamScore() int {
	offset := s.value % 10
	if offset >= 5 {
		return s.value/10 + 1
	}
	return s.value / 10
}

func NewScore(value int) Score {
	if value < 0 || value > 162 {
		panic("Invalid score value should be between 0 and 162")
	}
	return Score{value}
}
