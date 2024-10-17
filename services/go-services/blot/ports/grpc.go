package ports

import (
	"blot/internal/blot/app"
	"blot/internal/blot/app/command"
	"blot/internal/blot/domain/gameset"
	"blot/internal/blot/domain/team"
	blotservicepb "blot/internal/common/gen-proto/blotservice/v1beta1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	app app.Application
	blotservicepb.UnimplementedBlotServiceServer
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (g GrpcServer) CreateGameSet(ctx context.Context, req *blotservicepb.CreateGameSetRequest) (*blotservicepb.CreateGameSetResponse, error) {
	err := g.app.Commands.CreateGameSet.Handle(ctx, command.CreateGameSet{
		TeamID1:     team.NewID(req.Team1Id),
		TeamID2:     team.NewID(req.Team2Id),
		ID:          gameset.NewID(req.Id),
		FirstGameID: gameset.NewGameID(req.FirstGameId),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error()) // TODO: map error
	}
	return &blotservicepb.CreateGameSetResponse{}, nil
}

func (g GrpcServer) GetGameForPlayer(
	ctx context.Context,
	req *blotservicepb.GetGameForPlayerRequest,
) (*blotservicepb.GetGameForPlayerResponse, error) {
	return &blotservicepb.GetGameForPlayerResponse{
		CurrentPlayer: &blotservicepb.Player{
			Id:   "1",
			Name: "Player 1",
			HandCards: []*blotservicepb.Card{
				{
					Rank: blotservicepb.Rank_RANK_JACK,
					Suit: blotservicepb.Suit_SUIT_HEARTS,
				},
				{
					Rank: blotservicepb.Rank_RANK_JACK,
					Suit: blotservicepb.Suit_SUIT_DIAMONDS,
				},
				{
					Rank: blotservicepb.Rank_RANK_JACK,
					Suit: blotservicepb.Suit_SUIT_SPADES,
				},
				{
					Rank: blotservicepb.Rank_RANK_JACK,
					Suit: blotservicepb.Suit_SUIT_CLUBS,
				},
				{
					Rank: blotservicepb.Rank_RANK_QUEEN,
					Suit: blotservicepb.Suit_SUIT_HEARTS,
				},
				{
					Rank: blotservicepb.Rank_RANK_QUEEN,
					Suit: blotservicepb.Suit_SUIT_DIAMONDS,
				},
				{
					Rank: blotservicepb.Rank_RANK_QUEEN,
					Suit: blotservicepb.Suit_SUIT_SPADES,
				},
				//{
				//	value: blotservicepb.Rank_RANK_QUEEN,
				//	Suit: blotservicepb.Suit_SUIT_CLUBS,
				//},
			},
			DiscardStack: []*blotservicepb.Card{}, // TODO: Implement
			TeamId:       "1",
		},
		AllyPlayer: &blotservicepb.Player{
			Id:   "2",
			Name: "Player 2",
			HandCards: []*blotservicepb.Card{
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				//{
				//	value: blotservicepb.Rank_RANK_UNSPECIFIED,
				//	Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				//},
			},
			DiscardStack: []*blotservicepb.Card{}, // TODO: Implement
			TeamId:       "1",
		},
		LeftPlayer: &blotservicepb.Player{
			Id:   "3",
			Name: "Player 3",
			HandCards: []*blotservicepb.Card{
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				//{
				//	value: blotservicepb.Rank_RANK_UNSPECIFIED,
				//	Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				//},
			},
			DiscardStack: []*blotservicepb.Card{}, // TODO: Implement
			TeamId:       "2",
		},
		RightPlayer: &blotservicepb.Player{
			Id:   "4",
			Name: "Player 4",
			HandCards: []*blotservicepb.Card{
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				{
					Rank: blotservicepb.Rank_RANK_UNSPECIFIED,
					Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				},
				//{
				//	value: blotservicepb.Rank_RANK_UNSPECIFIED,
				//	Suit: blotservicepb.Suit_SUIT_UNSPECIFIED,
				//},
			},
			DiscardStack: []*blotservicepb.Card{}, // TODO: Implement
			TeamId:       "2",
		},
		Game: &blotservicepb.Game{
			Id:     "1",
			Status: blotservicepb.GameStatus_GAME_STATUS_CREATED,
			Round: &blotservicepb.Round{
				Number: 2,
				TableCards: []*blotservicepb.PlayerCard{
					{
						PlayerId: "2",
						Card:     &blotservicepb.Card{Rank: blotservicepb.Rank_RANK_ACE, Suit: blotservicepb.Suit_SUIT_HEARTS},
					},
					{
						PlayerId: "3",
						Card:     &blotservicepb.Card{Rank: blotservicepb.Rank_RANK_ACE, Suit: blotservicepb.Suit_SUIT_DIAMONDS},
					},
					//{
					//	PlayerId: "4",
					//	Card:     &blotservicepb.Card{value: blotservicepb.Rank_RANK_ACE, Suit: blotservicepb.Suit_SUIT_SPADES},
					//},
				},
				Status:          blotservicepb.RoundStatus_ROUND_STATUS_STARTED,
				CurrentPlayerId: "1",
			},
			Bet: &blotservicepb.Bet{
				Trump:  blotservicepb.Suit_SUIT_HEARTS,
				TeamId: "1",
				Amount: 8,
			},
			Teams: []*blotservicepb.Team{
				{
					Id:   "1",
					Name: "Team 1",
				},
				{
					Id:   "2",
					Name: "Team 2",
				},
			},
		},
	}, nil
}
