package adapters

import (
	"context"
	"sync"

	"blot/internal/blot/domain/gameset"
	"blot/internal/blot/domain/gameset/player"
)

type playerStorageModel struct {
	ID   string
	Name string
}

type gameSetStorageModel struct {
	ID      string
	OwnerID string
	players []playerStorageModel
	status  string
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
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.data[gameSet.ID().String()]; ok {
		return gameset.ErrGameSetAlreadyExists{ID: gameSet.ID()}
	}
	g.data[gameSet.ID().String()] = toGameSetStorageModel(gameSet)
	return nil
}

func (g *GameSetMemoryRepository) Get(ctx context.Context, id gameset.ID) (gameset.GameSet, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	if set, ok := g.data[id.String()]; ok {
		return toGameSet(set), nil
	}
	return gameset.GameSet{}, gameset.NotFoundError{ID: id}
}

func (g *GameSetMemoryRepository) GetByPlayerID(ctx context.Context, playerID player.ID) ([]gameset.GameSet, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	var res []gameset.GameSet
	for _, set := range g.data {
		for _, p := range set.players {
			if p.ID == playerID.String() {
				res = append(res, toGameSet(set))
			}
		}
	}
	return res, nil
}
func (g *GameSetMemoryRepository) UpdateByID(ctx context.Context, setID gameset.ID, updateFn func(set *gameset.GameSet) (bool, error)) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	setEntry, ok := g.data[setID.String()]
	if !ok {
		return gameset.NotFoundError{ID: setID}
	}
	set := toGameSet(setEntry)
	ok, err := updateFn(&set)
	if err != nil {
		return err
	}
	if !ok { // we don;t need to  update the set
		return nil
	}
	g.data[setID.String()] = toGameSetStorageModel(&set)
	return nil
}

func toGameSetStorageModel(set *gameset.GameSet) *gameSetStorageModel {
	return &gameSetStorageModel{
		ID:      set.ID().String(),
		OwnerID: set.OwnerID().String(),
		players: toPlayerStorageModel(set.Players()),
		status:  set.Status().String(),
	}
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
	return gameset.UnmarshalFromDatabase(id, gameset.NewGamesetStatus(set.status), firstPlayerID, toPlayers(set.players))
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
