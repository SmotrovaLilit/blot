package game

import (
	"errors"
	"strconv"
)

var RoundNumber1 = RoundNumber{1}

type RoundNumber struct {
	value int
}

const maxRoundNumber = 8

var ErrInvalidRoundNumber = errors.New("invalid rounds number")
var ErrLastRound = errors.New("it is last rounds")

func NewRoundNumber(value int) (RoundNumber, error) {
	if value < 1 || value > maxRoundNumber {
		return RoundNumber{}, ErrInvalidRoundNumber
	}
	return RoundNumber{value: value}, nil
}

func (n RoundNumber) String() string {
	return strconv.Itoa(n.value)
}

func (n RoundNumber) Next() (RoundNumber, error) {
	if n.value == maxRoundNumber {
		return RoundNumber{}, ErrLastRound
	}
	return RoundNumber{n.value + 1}, nil
}

func (n RoundNumber) Value() int {
	return n.value
}
