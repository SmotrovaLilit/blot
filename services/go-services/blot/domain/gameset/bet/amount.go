package bet

import "errors"

type Amount struct {
	value int
}

// TODO aSk.
const maxBetAmount = 50

var ErrInvalidAmount = errors.New("invalid amount")

func NewAmount(value int) (Amount, error) {
	if value < 0 || value > maxBetAmount {
		return Amount{}, ErrInvalidAmount
	}
	return Amount{value}, nil
}

func (a Amount) IsZero() bool {
	return Amount{} == a
}

func (a Amount) Value() int {
	return a.value
}

func MustNewAmount(i int) Amount {
	a, err := NewAmount(i)
	if err != nil {
		panic(err)
	}
	return a
}
