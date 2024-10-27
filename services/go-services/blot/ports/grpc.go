package ports

import (
	"context"
	"fmt"

	"blot/internal/blot/app/command/creategameset"

	"blot/internal/blot/app"
	"blot/internal/blot/app/command"
	"blot/internal/blot/app/query"
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset"
	"blot/internal/blot/domain/gameset/game"
	"blot/internal/blot/domain/gameset/player"
	"blot/internal/blot/domain/gameset/team"
	blotservicepb "blot/internal/common/gen-proto/blotservice/v1beta1"

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
	err := g.app.Commands.CreateGameSet.Handle(ctx, creategameset.CreateGameSet{
		ID:         req.Id,
		PlayerID:   req.PlayerId,
		PlayerName: req.PlayerName,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error()) // TODO: map error
	}
	return &blotservicepb.CreateGameSetResponse{}, nil
}

func (g GrpcServer) JoinGameSet(ctx context.Context, req *blotservicepb.JoinGameSetRequest) (*blotservicepb.JoinGameSetResponse, error) {
	err := g.app.Commands.JoinGameSet.Handle(ctx, command.JoinGameSet{
		ID:         req.Id,
		PlayerID:   req.PlayerId,
		PlayerName: req.PlayerName,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error()) // TODO: map error
	}
	return &blotservicepb.JoinGameSetResponse{}, nil
}

//func (g GrpcServer) LeaveGameSet(ctx context.Context, req *blotservicepb.LeaveGameSetRequest) (*blotservicepb.LeaveGameSetResponse, error) {
//	playerID, err := player.NewID(req.PlayerId)
//	if err != nil {
//		return nil, status.Error(codes.InvalidArgument, err.Error()) // TODO: map error
//	}
//	err = g.app.Commands.LeaveGameSet.Handle(ctx, command.LeaveGameSet{
//		ID:       gameset.NewID(req.Id),
//		PlayerID: playerID,
//	})
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error()) // TODO: map error
//	}
//	return &blotservicepb.LeaveGameSetResponse{}, nil
//}

func (g GrpcServer) StartGame(ctx context.Context, req *blotservicepb.StartGameRequest) (*blotservicepb.StartGameResponse, error) {
	err := g.app.Commands.StartGame.Handle(ctx, command.StartGame{
		SetID:    req.GameSetId,
		GameID:   req.GameId,
		PlayerID: req.PlayerId,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error()) // TODO: map error
	}
	return &blotservicepb.StartGameResponse{}, nil
}

func (g GrpcServer) GetGameSetForPlayer(ctx context.Context, req *blotservicepb.GetGameSetForPlayerRequest) (*blotservicepb.GetGameSetForPlayerResponse, error) {
	gameSet, err := g.app.Queries.GameSetForPlayer.Handle(ctx, query.GameSetForPlayer{
		GameSetID: req.Id,
		PlayerID:  req.PlayerId,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error()) // TODO: map error
	}
	r := gameSetToDetailResponse(*gameSet)
	return &blotservicepb.GetGameSetForPlayerResponse{
		GameSet: r,
	}, nil
}

func (g GrpcServer) GetGameSetsForPlayer(ctx context.Context, req *blotservicepb.GetGameSetsForPlayerRequest) (*blotservicepb.GetGameSetsForPlayerResponse, error) {
	gameSets, err := g.app.Queries.GameSetsForPlayer.Handle(ctx, query.GameSetsForPlayer{
		PlayerID: req.PlayerId,
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

func gameSetToDetailResponse(set gameset.GameSet) *blotservicepb.GameSet {
	return &blotservicepb.GameSet{
		Id:      set.ID().String(),
		OwnerId: set.OwnerID().String(),
		Players: playersToResponse(set.Players()),
		Status:  gameSetStatusToResponse(set.Status()),
		Game:    gameToResponse(set.LastGame()), // TODO: check if game exists
	}
}

func gameToResponse(lastGame game.Game) *blotservicepb.Game {
	if lastGame.IsZero() {
		return nil
	}
	// TODO try to us iterator
	playerStates := make([]*blotservicepb.PlayerStateInGame, len(lastGame.PlayerStates()))
	for i, ps := range lastGame.PlayerStates() {
		playerStates[i] = playerStateToResponse(ps)
	}
	return &blotservicepb.Game{
		Id:           lastGame.ID().String(),
		Status:       gameStatusToResponse(lastGame.Status()),
		Round:        nil,
		Bet:          nil,
		Team1:        teamToResponse(lastGame.FirstTeam()),
		Team2:        teamToResponse(lastGame.SecondTeam()),
		PlayerStates: playerStates,
	}
}

func playerStateToResponse(ps game.PlayerState) *blotservicepb.PlayerStateInGame {
	handCards := make([]*blotservicepb.Card, len(ps.HandCards()))
	for i, c := range ps.HandCards() {
		handCards[i] = &blotservicepb.Card{
			Suit: suitToResponse(c.Suit()),
			Rank: rankToResponse(c.Rank()),
		}
	}
	return &blotservicepb.PlayerStateInGame{
		Id:        ps.ID().String(),
		HandCards: handCards,
	}
}

func rankToResponse(rank card.Rank) blotservicepb.Rank {
	switch rank {
	case card.RankAce:
		return blotservicepb.Rank_RANK_ACE
	case card.RankKing:
		return blotservicepb.Rank_RANK_KING
	case card.RankQueen:
		return blotservicepb.Rank_RANK_QUEEN
	case card.RankJack:
		return blotservicepb.Rank_RANK_JACK
	case card.RankTen:
		return blotservicepb.Rank_RANK_TEN
	case card.RankNine:
		return blotservicepb.Rank_RANK_NINE
	case card.RankEight:
		return blotservicepb.Rank_RANK_EIGHT
	case card.RankSeven:
		return blotservicepb.Rank_RANK_SEVEN
	default:
		panic(fmt.Sprintf("unknown rank: %v", rank))
	}
}

func suitToResponse(suit card.Suit) blotservicepb.Suit {
	switch suit {
	case card.SuitClubs:
		return blotservicepb.Suit_SUIT_CLUBS
	case card.SuitDiamonds:
		return blotservicepb.Suit_SUIT_DIAMONDS
	case card.SuitHearts:
		return blotservicepb.Suit_SUIT_HEARTS
	case card.SuitSpades:
		return blotservicepb.Suit_SUIT_SPADES
	default:
		panic(fmt.Sprintf("unknown suit: %v", suit))
	}
}

func teamToResponse(team team.Team) *blotservicepb.Team {
	return &blotservicepb.Team{
		Player1: team.FirstPlayer().String(),
		Player2: team.SecondPlayer().String(),
	}
}

func gameStatusToResponse(s game.Status) blotservicepb.GameStatus {
	switch s {
	case game.GameStatusBetting:
		return blotservicepb.GameStatus_GAME_STATUS_BETTING
	case game.GameStatusPlaying:
		return blotservicepb.GameStatus_GAME_STATUS_PLAYING
	case game.GameStatusFinished:
		return blotservicepb.GameStatus_GAME_STATUS_FINISHED
	default:
		panic(fmt.Sprintf("unknown status: %v. Add new status in convert domain model to response function", s))
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

func gameSetStatusToResponse(status gameset.Status) blotservicepb.GameSetStatus {
	switch status {
	case gameset.StatusWaitedForPlayers:
		return blotservicepb.GameSetStatus_GAME_SET_STATUS_WAITED_FOR_PLAYERS
	case gameset.StatusReadyToStart:
		return blotservicepb.GameSetStatus_GAME_SET_STATUS_READY_TO_START
	case gameset.StatusPlaying:
		return blotservicepb.GameSetStatus_GAME_SET_STATUS_PLAYING
	default:
		panic(fmt.Sprintf("unknown status: %v. Add new status in convert domain model to response function", status))
	}
}
