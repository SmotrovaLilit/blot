package game

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset/player"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTableCards_CalculateWinner(t1 *testing.T) {
	type args struct {
		trump card.Suit
	}
	tests := []struct {
		name  string
		args  args
		table TableCards
		want  PlayerCard
	}{
		{
			name: "Case with no trump",
			table: NewTableCards([]PlayerCard{
				NewPlayerCard(player.MustNewID("1eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankKing, card.SuitSpades)),
				NewPlayerCard(player.MustNewID("2eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitSpades)),
				NewPlayerCard(player.MustNewID("3eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitDiamonds)),
				NewPlayerCard(player.MustNewID("4eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitClubs)),
			}),
			args: args{
				trump: card.SuitHearts,
			},
			want: NewPlayerCard(player.MustNewID("2eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitSpades)),
		},
		{
			name: "Case with trump",
			table: NewTableCards([]PlayerCard{
				NewPlayerCard(player.MustNewID("1eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankJack, card.SuitSpades)),
				NewPlayerCard(player.MustNewID("2eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitSpades)),
				NewPlayerCard(player.MustNewID("3eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitDiamonds)),
				NewPlayerCard(player.MustNewID("4eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitClubs)),
			}),
			args: args{
				trump: card.SuitSpades,
			},
			want: NewPlayerCard(player.MustNewID("1eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankJack, card.SuitSpades)),
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			got := tt.table.CalculateWinner(tt.args.trump)
			require.Equal(t1, tt.want, got)
		})
	}
}

func TestTableCards_CalculateScore(t1 *testing.T) {
	type args struct {
		trump card.Suit
	}
	tests := []struct {
		name  string
		args  args
		want  card.Score
		table TableCards
	}{
		{
			name: "Case with no trump",
			table: NewTableCards([]PlayerCard{
				NewPlayerCard(player.MustNewID("1eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankJack, card.SuitSpades)),
				NewPlayerCard(player.MustNewID("2eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitSpades)),
				NewPlayerCard(player.MustNewID("3eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitDiamonds)),
				NewPlayerCard(player.MustNewID("4eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitClubs)),
			}),
			args: args{
				trump: card.SuitHearts,
			},
			want: card.NewScore(35),
		},
		{
			name: "Case with trump",
			table: NewTableCards([]PlayerCard{
				NewPlayerCard(player.MustNewID("1eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankJack, card.SuitSpades)),
				NewPlayerCard(player.MustNewID("2eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitSpades)),
				NewPlayerCard(player.MustNewID("3eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitDiamonds)),
				NewPlayerCard(player.MustNewID("4eab4f96-2162-4d5b-b1ce-ea146f7736de"), card.NewCard(card.RankAce, card.SuitClubs)),
			}),
			args: args{
				trump: card.SuitSpades,
			},
			want: card.NewScore(53),
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			got := tt.table.CalculateScore(tt.args.trump)
			require.Equal(t1, tt.want, got)
		})
	}
}
