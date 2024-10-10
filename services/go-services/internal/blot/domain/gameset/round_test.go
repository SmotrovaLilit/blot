package gameset

//
//import (
//	"github.com/stretchr/testify/require"
//	"blot/internal/blot/domain/user"
//	"testing"
//)
//
//func TestRound_Turn(t *testing.T) {
//	u1 := CreatePlayerInTeam1(user.NewUser(user.NewID("1f40f3d5-8964-4cdb-baac-3d9d0c25b80b"), "user1"))
//	u2 := CreatePlayerInTeam2(user.NewUser(user.NewID("2f40f3d5-8964-4cdb-baac-3d9d0c25b80b"), "user2"))
//	u3 := CreatePlayerInTeam1(user.NewUser(user.NewID("3f40f3d5-8964-4cdb-baac-3d9d0c25b80b"), "user3"))
//	u4 := CreatePlayerInTeam2(user.NewUser(user.NewID("4f40f3d5-8964-4cdb-baac-3d9d0c25b80b"), "user4"))
//	type fields struct {
//		number RoundNumber
//	}
//	type args struct {
//		sittingOrder SittingOrder
//	}k
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   Player
//	}{
//		{
//			name: "1 round",
//			fields: fields{
//				number: NewRoundNumber(1),
//			},
//			args: args{
//				sittingOrder: NewPlayersSittingOrder(u1, u2, u3, u4),
//			},
//			want: u1,
//		},
//		{
//			name: "2 round",
//			fields: fields{
//				number: NewRoundNumber(2),
//			},
//			args: args{
//				sittingOrder: NewPlayersSittingOrder(u1, u2, u3, u4),
//			},
//			want: u2,
//		},
//		{
//			name: "3 round",
//			fields: fields{
//				number: NewRoundNumber(3),
//			},
//			args: args{
//				sittingOrder: NewPlayersSittingOrder(u1, u2, u3, u4),
//			},
//			want: u3,
//		},
//		{
//			name: "4 round",
//			fields: fields{
//				number: NewRoundNumber(4),
//			},
//			args: args{
//				sittingOrder: NewPlayersSittingOrder(u1, u2, u3, u4),
//			},
//			want: u4,
//		},
//		{
//			name: "5 round",
//			fields: fields{
//				number: NewRoundNumber(5),
//			},
//			args: args{
//				sittingOrder: NewPlayersSittingOrder(u1, u2, u3, u4),
//			},
//			want: u1,
//		},
//		{
//			name: "8 round",
//			fields: fields{
//				number: NewRoundNumber(8),
//			},
//			args: args{
//				sittingOrder: NewPlayersSittingOrder(u1, u2, u3, u4),
//			},
//			want: u4,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := Round{
//				number: tt.fields.number,
//			}
//			got := r.firstTurnIndex(tt.args.sittingOrder)
//			require.Equal(t, tt.want, got)
//		})
//	}
//}
