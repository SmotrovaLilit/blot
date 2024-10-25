package gameset

import "errors"

var (
	ErrGameSetNotReadyToStartGame = errors.New("game set is not ready to start the game")

	StatusWaitedForPlayers = Status{"waited_for_players"}
	StatusReadyToStart     = Status{"ready_to_start"}
	StatusPlaying          = Status{"playing"}
	Statuses               = []Status{StatusWaitedForPlayers, StatusReadyToStart, StatusPlaying}
)

type Status struct {
	value string
}

func NewStatus(value string) Status {
	for _, s := range Statuses {
		if s.value == value {
			return s
		}
	}
	panic("invalid status")
}

func (s Status) String() string {
	return s.value
}

func (s Status) CanJoin() bool {
	return s == StatusWaitedForPlayers
}

func (s Status) StartGame() (Status, error) {
	if s != StatusReadyToStart {
		return Status{}, ErrGameSetNotReadyToStartGame
	}

	return StatusPlaying, nil

}
