package adapters

import (
	"context"
	"log/slog"
	"sync"

	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset/game"
	"blot/internal/blot/domain/gameset/team"
	"blot/internal/common/logging"

	"go.opentelemetry.io/otel"

	"blot/internal/blot/domain/gameset"
	"blot/internal/blot/domain/gameset/player"
)

var tracer = otel.Tracer("repository")

type playerStorageModel struct {
	ID   string
	Name string
}

type gameSetStorageModel struct {
	ID      string
	OwnerID string
	Players []playerStorageModel
	Status  string
	Game    *gameStorageModel
}

type playerStateStorageModel struct {
	ID        string
	HandCards []storageCard
}

type storageCard struct {
	rank string
	suit string
}

type teamStorageModel struct {
	ID      string
	Players [2]string
}

type gameStorageModel struct {
	ID      string
	Status  string
	Players []playerStateStorageModel
	team1   teamStorageModel
	team2   teamStorageModel
}

type GameSetMemoryRepository struct {
	data map[string]*gameSetStorageModel
	mu   sync.RWMutex
}

func NewGameSetMemoryRepository() *GameSetMemoryRepository {
	return &GameSetMemoryRepository{
		data: make(map[string]*gameSetStorageModel),
	}
}

func (g *GameSetMemoryRepository) Create(ctx context.Context, gameSet *gameset.GameSet) error {
	ctx, span := tracer.Start(ctx, "gamSetRepo.Create")
	defer span.End()
	ctx = logging.AppendCtx(ctx, slog.String("repo_method", "Create"), slog.Any("set", *gameSet))
	slog.DebugContext(ctx, "repo: creating game set")
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.data[gameSet.ID().String()]; ok {
		slog.ErrorContext(ctx, "repo: failed to create game set: already exists")
		return gameset.ErrGameSetAlreadyExists{ID: gameSet.ID()}
	}
	g.data[gameSet.ID().String()] = toGameSetStorageModel(gameSet)
	slog.DebugContext(ctx, "repo: game set created")
	return nil
}

func (g *GameSetMemoryRepository) Get(ctx context.Context, id gameset.ID) (gameset.GameSet, error) {
	ctx, span := tracer.Start(ctx, "gamSetRepo.Get")
	defer span.End()
	ctx = logging.AppendCtx(ctx, slog.String("repo_method", "Get"), slog.String("id", id.String()))
	slog.DebugContext(ctx, "repo: getting game setEntry")
	g.mu.RLock()
	defer g.mu.RUnlock()
	if setEntry, ok := g.data[id.String()]; ok {
		existGame := toGameSet(setEntry)
		slog.DebugContext(
			ctx,
			"repo: game setEntry found",
			slog.Any("set", existGame),
			slog.Any("set_entry", setEntry),
		)
		return existGame, nil
	}
	slog.DebugContext(ctx, "repo: game setEntry not found")
	return gameset.GameSet{}, gameset.NotFoundError{ID: id}
}

func (g *GameSetMemoryRepository) GetByPlayerID(ctx context.Context, playerID player.ID) ([]gameset.GameSet, error) {
	ctx, span := tracer.Start(ctx, "gamSetRepo.GetByPlayerID")
	defer span.End()
	g.mu.RLock()
	ctx = logging.AppendCtx(ctx, slog.String("repo_method", "GetByPlayerID"), slog.String("player_id", playerID.String()))
	defer g.mu.RUnlock()
	var res []gameset.GameSet
	for _, set := range g.data {
		for _, p := range set.Players {
			if p.ID == playerID.String() {
				res = append(res, toGameSet(set))
			}
		}
	}
	slog.DebugContext(ctx, "repo: get player's game set succeed", slog.Any("game_sets", res))
	return res, nil
}
func (g *GameSetMemoryRepository) UpdateByID(ctx context.Context, setID gameset.ID, updateFn func(set *gameset.GameSet) (bool, error)) error {
	ctx, span := tracer.Start(ctx, "gamSetRepo.UpdateByID")
	defer span.End()
	ctx = logging.AppendCtx(ctx, slog.String("repo_method", "UpdateByID"), slog.String("id", setID.String()))
	slog.DebugContext(ctx, "repo: updating game set")
	g.mu.Lock()
	defer g.mu.Unlock()
	setEntry, ok := g.data[setID.String()]
	if !ok {
		return gameset.NotFoundError{ID: setID}
	}
	set := toGameSet(setEntry)
	ctx = logging.AppendCtx(ctx, slog.Any("set", set), slog.Any("set_entry", setEntry))
	slog.DebugContext(ctx, "repo: got game set to update from memory")
	ok, err := updateFn(&set)
	if err != nil {
		slog.ErrorContext(ctx, "repo: updateFn returns error. set not updated", slog.Any("error", err))
		return err
	}
	if !ok { // we don;t need to  update the set
		slog.DebugContext(ctx, "repo: updateFn returns false, set not updated")
		return nil
	}
	setEntry = toGameSetStorageModel(&set)
	g.data[setID.String()] = setEntry
	slog.DebugContext(ctx, "repo: set updated", slog.Any("set_updated", set), slog.Any("set_entry_updated", setEntry))
	return nil
}

func toGameSetStorageModel(set *gameset.GameSet) *gameSetStorageModel {
	return &gameSetStorageModel{
		ID:      set.ID().String(),
		OwnerID: set.OwnerID().String(),
		Players: toPlayerStorageModel(set.Players()),
		Status:  set.Status().String(),
		Game:    toGameStorageModel(set.LastGame()),
	}
}

func toGameStorageModel(game game.Game) *gameStorageModel {
	if game.IsZero() {
		return nil
	}
	return &gameStorageModel{
		ID:      game.ID().String(),
		Status:  game.Status().String(),
		Players: toPlayerStateStorageModel(game.PlayerStates()),
		team1:   toTeamStorageModel(game.FirstTeam()),
		team2:   toTeamStorageModel(game.SecondTeam()),
	}
}

func toTeamStorageModel(team team.Team) teamStorageModel {
	return teamStorageModel{
		ID:      team.ID().String(),
		Players: [2]string{team.FirstPlayer().String(), team.SecondPlayer().String()},
	}
}

func toPlayerStateStorageModel(states []game.PlayerState) []playerStateStorageModel {
	res := make([]playerStateStorageModel, len(states))
	for i, s := range states {
		sCards := make([]storageCard, len(s.HandCards()))
		for j, c := range s.HandCards() {
			sCards[j] = storageCard{
				rank: c.Rank().String(),
				suit: c.Suit().String(),
			}
		}
		res[i] = playerStateStorageModel{
			ID:        s.ID().String(),
			HandCards: sCards,
		}
	}
	return res
}

func toPlayerStorageModel(players []player.Player) []playerStorageModel {
	var res []playerStorageModel
	for _, p := range players {
		res = append(res, playerStorageModel{
			ID:   p.ID().String(),
			Name: p.Name().String(),
		})
	}
	return res
}

func toGameSet(set *gameSetStorageModel) gameset.GameSet {
	firstPlayerID, err := player.NewID(set.OwnerID)
	if err != nil {
		panic(err)
	}
	id := gameset.NewID(set.ID)
	return gameset.UnmarshalFromDatabase(id, gameset.NewStatus(set.Status), firstPlayerID, toPlayers(set.Players), toGame(set.Game))
}

func toGame(model *gameStorageModel) game.Game {
	if model == nil {
		return game.Game{}
	}
	id := game.NewID(model.ID)
	team1 := toTeam(model.team1)
	team2 := toTeam(model.team2)
	return game.UnmarshalFromDatabase(id, game.NewStatus(model.Status), team1, team2, toPlayerStates(model.Players))
}

func toPlayerStates(players []playerStateStorageModel) []game.PlayerState {
	var res []game.PlayerState
	for _, p := range players {
		id, err := player.NewID(p.ID)
		if err != nil {
			panic(err)
		}
		res = append(res, game.UnmarshalFromDatabasePlayerState(id, toCards(p.HandCards)))
	}
	return res
}

func toCards(cards []storageCard) []card.Card {
	var res []card.Card
	for _, c := range cards {
		res = append(res, card.UnmarshalFromDatabase(card.NewRank(c.rank), card.NewSuit(c.suit)))
	}
	return res
}

func toTeam(team1 teamStorageModel) team.Team {
	id := team.NewID(team1.ID)
	p1, err := player.NewID(team1.Players[0])
	if err != nil {
		panic(err)
	}
	p2, err := player.NewID(team1.Players[1])
	if err != nil {
		panic(err)
	}

	return team.UnmarshalFromDatabase(id, p1, p2)
}

func toPlayers(players []playerStorageModel) []player.Player {
	var res []player.Player
	for _, p := range players {
		domainPlayer, err := player.Create(p.ID, p.Name)
		if err != nil {
			panic(err)
		}
		res = append(res, domainPlayer)
	}
	return res
}
