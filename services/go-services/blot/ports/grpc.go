package ports

import (
	"blot/internal/blot/app"
	"blot/internal/blot/app/command"
	"blot/internal/blot/app/query"
	"blot/internal/blot/domain/gameset"
	"blot/internal/blot/domain/gameset/player"
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
	pl, err := player.Create(req.PlayerId, req.PlayerName)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error()) // TODO: map error
	}
	err = g.app.Commands.CreateGameSet.Handle(ctx, command.CreateGameSet{
		ID:          gameset.NewID(req.Id),
		FirstPlayer: pl,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error()) // TODO: map error
	}
	return &blotservicepb.CreateGameSetResponse{}, nil
}

func (g GrpcServer) JoinGameSet(ctx context.Context, req *blotservicepb.JoinGameSetRequest) (*blotservicepb.JoinGameSetResponse, error) {
	pl, err := player.Create(req.PlayerId, req.PlayerName)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error()) // TODO: map error
	}
	err = g.app.Commands.JoinGameSet.Handle(ctx, command.JoinGameSet{
		ID:     gameset.NewID(req.Id),
		Player: pl,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error()) // TODO: map error
	}
	return &blotservicepb.JoinGameSetResponse{}, nil
}

func (g GrpcServer) LeaveGameSet(ctx context.Context, req *blotservicepb.LeaveGameSetRequest) (*blotservicepb.LeaveGameSetResponse, error) {
	playerID, err := player.NewID(req.PlayerId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error()) // TODO: map error
	}
	err = g.app.Commands.LeaveGameSet.Handle(ctx, command.LeaveGameSet{
		ID:       gameset.NewID(req.Id),
		PlayerID: playerID,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error()) // TODO: map error
	}
	return &blotservicepb.LeaveGameSetResponse{}, nil
}

func (g GrpcServer) GetGameSetForPlayer(ctx context.Context, req *blotservicepb.GetGameSetForPlayerRequest) (*blotservicepb.GetGameSetForPlayerResponse, error) {
	playerID, err := player.NewID(req.PlayerId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error()) // TODO: map error
	}
	gameSet, err := g.app.Queries.GameSetForPlayer.Handle(ctx, query.GameSetForPlayer{
		GameSetID: gameset.NewID(req.Id),
		PlayerID:  playerID,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error()) // TODO: map error
	}
	r := gameSetToResponse(*gameSet)
	return &blotservicepb.GetGameSetForPlayerResponse{
		GameSet: r,
	}, nil
}

func (g GrpcServer) GetGameSetsForPlayer(ctx context.Context, req *blotservicepb.GetGameSetsForPlayerRequest) (*blotservicepb.GetGameSetsForPlayerResponse, error) {
	playerID, err := player.NewID(req.PlayerId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error()) // TODO: map error
	}
	gameSets, err := g.app.Queries.GameSetsForPlayer.Handle(ctx, query.GameSetsForPlayer{
		PlayerID: playerID,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error()) // TODO: map error
	}
	var res []*blotservicepb.GameSet
	for _, set := range gameSets {
		res = append(res, gameSetToResponse(set))
	}
	return &blotservicepb.GetGameSetsForPlayerResponse{
		GameSets: res,
	}, nil
}

func gameSetToResponse(set gameset.GameSet) *blotservicepb.GameSet {
	return &blotservicepb.GameSet{
		Id:      set.ID().String(),
		OwnerId: set.OwnerID().String(),
		Status:  gameSetStatusToResponse(set.Status()),
		Players: playersToResponse(set.Players()),
	}
}

func playersToResponse(players []player.Player) []*blotservicepb.Player {
	var res []*blotservicepb.Player
	for _, p := range players {
		res = append(res, &blotservicepb.Player{
			Id:   p.ID().String(),
			Name: p.Name().String(),
		})
	}
	return res
}

func gameSetStatusToResponse(status gameset.GamesetStatus) blotservicepb.GameSetStatus {
	switch status {
	case gameset.GamesetStatusWaitedForPlayers:
		return blotservicepb.GameSetStatus_GAME_SET_STATUS_WAITED_FOR_PLAYERS
	case gameset.GamesetStatusReadyToStart:
		return blotservicepb.GameSetStatus_GAME_SET_STATUS_READY_TO_START
	default:
		return blotservicepb.GameSetStatus_GAME_SET_STATUS_UNSPECIFIED
	}
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
