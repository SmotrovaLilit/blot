package gameset

import (
	"fmt"
)

var (
	StatusWaitedForPlayers = Status{"waited_for_players"}
	StatusReadyToStart     = Status{"ready_to_start"}
	StatusPlaying          = Status{"playing"}
	Statuses               = []Status{StatusWaitedForPlayers, StatusReadyToStart, StatusPlaying}
)

type ErrGameSetNotReadyToStartGame struct {
	Status string
}

func (e ErrGameSetNotReadyToStartGame) Error() string {
	return fmt.Sprintf("game set is not ready to start the game: %s", e.Status)
}

type ErrGameSetNotReadyToPlayCard struct {
	Status string
}

func (e ErrGameSetNotReadyToPlayCard) Error() string {
	return fmt.Sprintf("game set is not ready to play card, current status: %s", e.Status)
}

type ErrGameSetNotReadyToSetBet struct {
	Status string
}

func (e ErrGameSetNotReadyToSetBet) Error() string {
	return fmt.Sprintf("game set is not ready to set bet, current status: %s", e.Status)
}

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
		return Status{}, ErrGameSetNotReadyToStartGame{
			Status: s.String(),
		}
	}

	return StatusPlaying, nil
}

func (s Status) IsZero() bool {
	return s == Status{}
}

func (s Status) CanPlayCard() bool {
	return s == StatusPlaying
}

func (s Status) CanSetBet() bool {
	return s == StatusPlaying
}
