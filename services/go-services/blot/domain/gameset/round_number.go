package gameset

const (
	MaxRoundNumber = 8
)

type RoundNumber struct {
	number int
}

func (n RoundNumber) IsLastRound() bool {
	return n.number == MaxRoundNumber
}

func (n RoundNumber) Next() RoundNumber {
	if n.IsLastRound() {
		panic("last round")
	}
	return RoundNumber{n.number + 1}
}

func (n RoundNumber) Number() int {
	return n.number
}

func NewRoundNumber(number int) RoundNumber {
	if number < 0 || number > MaxRoundNumber {
		panic("invalid round number")
	}
	return RoundNumber{number}
}
