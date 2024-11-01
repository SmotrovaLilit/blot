package card

import (
	"fmt"
)

var (
	RankSeven = Rank{"seven", NewScore(0), NewScore(0), NewScore(0)}
	RankEight = Rank{"eight", NewScore(0), NewScore(0), NewScore(0)}
	RankNine  = Rank{"nine", NewScore(0), NewScore(19), NewScore(0)}
	RankTen   = Rank{"ten", NewScore(10), NewScore(10), NewScore(10)}
	RankJack  = Rank{"jack", NewScore(2), NewScore(20), NewScore(2)}
	RankQueen = Rank{"queen", NewScore(3), NewScore(3), NewScore(3)}
	RankKing  = Rank{"king", NewScore(4), NewScore(4), NewScore(4)}
	RankAce   = Rank{"ace", NewScore(11), NewScore(11), NewScore(19)}

	Ranks = []Rank{RankSeven, RankEight, RankNine, RankTen, RankJack, RankQueen, RankKing, RankAce}
)

type Rank struct {
	value        string
	score        Score
	scoreIfTrump Score
	scoreIfBoi   Score
}

func (r Rank) GetTrumpScore() Score {
	return r.scoreIfTrump
}

func (r Rank) GetScore() Score {
	return r.score
}

func (r Rank) String() string {
	return r.value
}

func (r Rank) IsZero() bool {
	return r == Rank{}
}

func NewRank(rankString string) Rank {
	for _, Rank := range Ranks {
		if Rank.value == rankString {
			return Rank
		}
	}
	panic(fmt.Sprintf("Invalid rank: %s", rankString))
}
