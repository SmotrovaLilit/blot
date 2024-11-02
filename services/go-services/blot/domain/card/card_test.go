package card

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCard_Beats(t *testing.T) {
	type args struct {
		winnerCard Card
		trump      Suit
	}
	tests := []struct {
		name string
		args args
		card Card
		want bool
	}{
		{
			name: "Ace of Spades beats 10 of Spades. Trump is something else",
			card: NewCard(RankAce, SuitSpades),
			args: args{
				winnerCard: NewCard(RankTen, SuitSpades),
				trump:      SuitDiamonds,
			},
			want: true,
		},
		{
			name: "10 of Spades beats King of Spades. Trump is something else",
			card: NewCard(RankTen, SuitSpades),
			args: args{
				winnerCard: NewCard(RankKing, SuitSpades),
				trump:      SuitDiamonds,
			},
			want: true,
		},
		{
			name: "King of Diamonds beats Queen of Diamonds. Trump is something else",
			card: NewCard(RankKing, SuitDiamonds),
			args: args{
				winnerCard: NewCard(RankQueen, SuitDiamonds),
				trump:      SuitSpades,
			},
			want: true,
		},
		{
			name: "Queen of Diamonds beats Jack of Diamonds. Trump is something else",
			card: NewCard(RankQueen, SuitDiamonds),
			args: args{
				winnerCard: NewCard(RankJack, SuitDiamonds),
				trump:      SuitSpades,
			},
			want: true,
		},
		{
			name: "Jack of Diamonds beats 9 of Diamonds. Trump is something else",
			card: NewCard(RankJack, SuitDiamonds),
			args: args{
				winnerCard: NewCard(RankNine, SuitDiamonds),
				trump:      SuitSpades,
			},
			want: true,
		},
		{
			name: "Jack of Diamonds doesn't beat 9 of Spades. Trump is something else",
			card: NewCard(RankJack, SuitDiamonds),
			args: args{
				winnerCard: NewCard(RankNine, SuitSpades),
				trump:      SuitClubs,
			},
			want: false,
		},
		{
			name: "Jack of Diamonds(trump) beats Ace of Spades. Trump is Diamonds",
			card: NewCard(RankJack, SuitDiamonds),
			args: args{
				winnerCard: NewCard(RankAce, SuitSpades),
				trump:      SuitDiamonds,
			},
			want: true,
		},
		{
			name: "Nine of Diamonds(trump) beats Ace of Spades. Trump is Diamonds",
			card: NewCard(RankNine, SuitDiamonds),
			args: args{
				winnerCard: NewCard(RankAce, SuitSpades),
				trump:      SuitDiamonds,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.card.Beats(tt.args.winnerCard, tt.args.trump)
			require.Equal(t, tt.want, got)
		})
	}
}
