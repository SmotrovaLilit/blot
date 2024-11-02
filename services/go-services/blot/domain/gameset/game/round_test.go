package game

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset/player"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRound_PlayCard(t *testing.T) {
	type args struct {
		c PlayerCard
	}
	tests := []struct {
		name          string
		round         Round
		args          args
		shouldFail    bool
		expectedError error
	}{
		{
			name: "should play card in first turn",
			round: Round{
				number: RoundNumber{value: 1},
				table:  NewTableCards([]PlayerCard{}),
			},
			args: args{
				c: NewPlayerCard(player.MustNewID("1eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitSpades)),
			},
		},
		{
			name: "should play card in last turn",
			round: Round{
				number: RoundNumber{value: 1},
				table: NewTableCards([]PlayerCard{
					NewPlayerCard(player.MustNewID("1eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitSpades)),
					NewPlayerCard(player.MustNewID("2eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitSpades)),
					NewPlayerCard(player.MustNewID("3eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitHearts)),
				}),
			},
			args: args{
				c: NewPlayerCard(player.MustNewID("5eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitSpades)),
			},
		},
		{
			name: "should fail when rounds is full",
			round: Round{
				number: RoundNumber{value: 1},
				table: NewTableCards([]PlayerCard{
					NewPlayerCard(player.MustNewID("1eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitSpades)),
					NewPlayerCard(player.MustNewID("2eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitHearts)),
					NewPlayerCard(player.MustNewID("3eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitClubs)),
					NewPlayerCard(player.MustNewID("4eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitDiamonds)),
				}),
			},
			args: args{
				c: NewPlayerCard(player.MustNewID("5eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankKing, card.SuitSpades)),
			},
			shouldFail:    true,
			expectedError: ErrTableFull,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newRound, err := tt.round.PlayCard(tt.args.c)
			if tt.shouldFail {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.args.c, newRound.table.cards[len(newRound.table.cards)-1])
		})
	}
}
