package gameset

const MaxTurnsCount = 4

type Round struct {
	number RoundNumber // 1-8
	turns  []Turn
}

func NewRound(number RoundNumber, turns []Turn) Round {
	if len(turns) > MaxTurnsCount {
		panic("too many turns")
	}
	return Round{number: number, turns: turns}
}

// Func (r Round) PlayCard(p Player, c card.card) {
//	if r.Finished() {
//		panic("round is finished")
//	}
//	if !r.IsPlayerTurn(p) {
//		panic("not your turn")
//	}
//	p.PlayCard(c)
//	//r.turns = append(r.turns, NewPlayedCard(c, p))
// }.

func (r Round) Number() RoundNumber {
	return r.number
}

func (r Round) IsLastRound() bool {
	return r.number.IsLastRound()
}

func (r Round) Finished() bool {
	return len(r.turns) == MaxTurnsCount
}

// Func (r Round) Winner(trump card.Suit) Player {
//	if !r.Finished() {
//		panic("round is not finished")
//	}
//	//winCardNumber := 0
//	//for i := 1; i < len(r.playedCards); i++ {
//	//	if r.playedCards[i].card.IsStronger(r.playedCards[winCardNumber].card, trump) {
//	//		winCardNumber = i
//	//	}
//	//}
//	//return r.playedCards[winCardNumber].Player
//	return Player{}
// }.

// Func (r Round) IsPlayerTurn(player Player) bool {
//	if r.Finished() {
//		panic("round is finished")
//	}
//	return r.currentTurnPlayer().Equal(player)
//}
//
// func (r Round) currentTurnPlayer() Player {
//	if r.Finished() {
//		panic("round is finished")
//	}
//	//return r.sittingOrder[r.currentTurnIndex()]
//	return Player{}
// }.

// firstTurnIndex can return 0, 1, 2, 3. It is the index of the current player in the round.
func (r Round) currentTurnIndex() int {
	if r.Finished() {
		panic("round is finished")
	}
	// playedCardsCount := len(r.playedCards)
	// return (r.firstTurnIndex() + playedCardsCount) % MaxCardsCount
	return 0
}

// firstTurnIndex can return 0, 1, 2, 3. It is the index of the first player in the round.
func (r Round) firstTurnIndex() int {
	// return (r.number.Number() - 1) % MaxCardsCount
	return 0
}
