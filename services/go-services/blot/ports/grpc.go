package ports

import (
	"context"
	"fmt"

	"blot/internal/blot/domain/gameset/game/bet"

	"blot/internal/blot/app/command/playcard"
	"blot/internal/blot/app/command/setbet"

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

// Func (g GrpcServer) LeaveGameSet(ctx context.Context, req *blotservicepb.LeaveGameSetRequest) (*blotservicepb.LeaveGameSetResponse, error) {
//	playerID, err := player.NewID(req.PlayerId)
//	if err != nil {
//		return nil, status.Error(codes.InvalidArgument, err.Error()) // TODO: map error
//	}
//	err = g.app.Commands.LeaveGameSet.Handle(ctx, command.LeaveGameSet{
//		ID:       gameset.NewID(req.Id),
//		playerID: playerID,
//	})
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error()) // TODO: map error
//	}
//	return &blotservicepb.LeaveGameSetResponse{}, nil
// }.

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

func (g GrpcServer) SetBet(ctx context.Context, req *blotservicepb.SetBetRequest) (*blotservicepb.SetBetResponse, error) {
	setID, err := gameset.NewID(req.GameSetId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error()) // TODO: map error
	}
	playerID, err := player.NewID(req.PlayerId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error()) // TODO: map error
	}
	suit, err := toSuit(req.Trump)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error()) // TODO: map error
	}
	am, err := bet.NewAmount(int(req.Amount))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error()) // TODO: map error
	}
	err = g.app.Commands.SetBet.Handle(ctx, setbet.SetBet{
		SetID:     setID,
		PlayerID:  playerID,
		BetTrump:  suit,
		BetAmount: am,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error()) // TODO: map error
	}
	return &blotservicepb.SetBetResponse{}, nil
}

func (g GrpcServer) PlayCard(ctx context.Context, req *blotservicepb.PlayCardRequest) (*blotservicepb.PlayCardResponse, error) {
	id, err := gameset.NewID(req.GameSetId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error()) // TODO: map error
	}
	playerID, err := player.NewID(req.PlayerId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error()) // TODO: map error
	}
	c, err := toCard(req.Card)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error()) // TODO: map error
	}
	err = g.app.Commands.PlayCard.Handle(ctx, playcard.PlayCard{
		SetID:    id,
		PlayerID: playerID,
		Card:     c,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error()) // TODO: map error
	}
	return &blotservicepb.PlayCardResponse{}, nil
}

func toCard(c *blotservicepb.Card) (card.Card, error) {
	rank, err := toRank(c.Rank)
	if err != nil {
		return card.Card{}, err
	}
	suit, err := toSuit(c.Suit)
	if err != nil {
		return card.Card{}, err
	}
	return card.NewCard(rank, suit), nil
}

func toSuit(suit blotservicepb.Suit) (card.Suit, error) {
	switch suit {
	case blotservicepb.Suit_SUIT_CLUBS:
		return card.SuitClubs, nil
	case blotservicepb.Suit_SUIT_DIAMONDS:
		return card.SuitDiamonds, nil
	case blotservicepb.Suit_SUIT_HEARTS:
		return card.SuitHearts, nil
	case blotservicepb.Suit_SUIT_SPADES:
		return card.SuitSpades, nil
	case blotservicepb.Suit_SUIT_UNSPECIFIED:
		return card.Suit{}, fmt.Errorf("unspecified suit")
	default:
		return card.Suit{}, fmt.Errorf("unknown suit: %v", suit)
	}
}

func toRank(rank blotservicepb.Rank) (card.Rank, error) {
	switch rank {
	case blotservicepb.Rank_RANK_ACE:
		return card.RankAce, nil
	case blotservicepb.Rank_RANK_KING:
		return card.RankKing, nil
	case blotservicepb.Rank_RANK_QUEEN:
		return card.RankQueen, nil
	case blotservicepb.Rank_RANK_JACK:
		return card.RankJack, nil
	case blotservicepb.Rank_RANK_TEN:
		return card.RankTen, nil
	case blotservicepb.Rank_RANK_NINE:
		return card.RankNine, nil
	case blotservicepb.Rank_RANK_EIGHT:
		return card.RankEight, nil
	case blotservicepb.Rank_RANK_SEVEN:
		return card.RankSeven, nil
	case blotservicepb.Rank_RANK_UNSPECIFIED:
		return card.Rank{}, fmt.Errorf("unspecified rank")
	default:
		return card.Rank{}, fmt.Errorf("unknown rank: %v", rank)
	}
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
	domainPlayers := lastGame.PlayerStates()
	playerStates := make([]*blotservicepb.PlayerStateInGame, len(domainPlayers))
	for i, ps := range domainPlayers {
		playerStates[i] = playerStateToResponse(ps)
	}
	domainRounds := lastGame.Rounds()
	rounds := make([]*blotservicepb.Round, len(domainRounds))
	for i, r := range domainRounds {
		rounds[i] = roundToResponse(r, lastGame.Bet().Trump())
	}
	return &blotservicepb.Game{
		Id:                  lastGame.ID().String(),
		Status:              gameStatusToResponse(lastGame.Status()),
		Rounds:              rounds,
		Bet:                 betToResponse(lastGame.Bet()),
		Team1:               teamToResponse(lastGame.FirstTeam()),
		Team2:               teamToResponse(lastGame.SecondTeam()),
		PlayerStates:        playerStates,
		CurrentTurnPlayerId: lastGame.CurrentTurnPlayerID().String(),
	}
}

func roundToResponse(r game.Round, trump card.Suit) *blotservicepb.Round {
	domainTableCards := r.TableCards()
	cards := make([]*blotservicepb.PlayedCard, len(domainTableCards))
	for i, c := range domainTableCards {
		cards[i] = toPlayerCardResponse(c)
	}

	return &blotservicepb.Round{
		// nolint:gosec
		Number:     int32(r.Number().Value()),
		TableCards: cards,
		WinnerId:   r.CalculateWinner(trump).PlayerID().String(),
		// nolint:gosec
		Score: int32(r.CalculateScore(trump).Value()),
	}
}

func toPlayerCardResponse(c game.PlayerCard) *blotservicepb.PlayedCard {
	return &blotservicepb.PlayedCard{
		PlayerId: c.PlayerID().String(),
		Card: &blotservicepb.Card{
			Suit: suitToResponse(c.Card().Suit()),
			Rank: rankToResponse(c.Card().Rank()),
		},
	}
}

func betToResponse(b bet.Bet) *blotservicepb.Bet {
	if b.IsZero() {
		return nil
	}
	return &blotservicepb.Bet{
		Trump: suitToResponse(b.Trump()),
		// nolint:gosec
		Amount: int32(b.Amount().Value()),
		TeamId: b.TeamID().String(),
	}
}

func playerStateToResponse(ps game.PlayerState) *blotservicepb.PlayerStateInGame {
	domainHandCards := ps.HandCards()
	handCards := make([]*blotservicepb.Card, len(domainHandCards))
	for i, c := range domainHandCards {
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
		Id:      team.ID().String(),
		Player1: team.FirstPlayer().String(),
		Player2: team.SecondPlayer().String(),
	}
}

func gameStatusToResponse(s game.Status) blotservicepb.GameStatus {
	switch s {
	case game.StatusBetting:
		return blotservicepb.GameStatus_GAME_STATUS_BETTING
	case game.StatusPlaying:
		return blotservicepb.GameStatus_GAME_STATUS_PLAYING
	case game.StatusFinished:
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
