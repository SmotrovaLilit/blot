package gameset

type GameStatus struct {
	value string
}

func (s GameStatus) IsFinished() bool {
	return s == GameStatusFinished
}

func (s GameStatus) CanPlayCard() bool {
	return s == GameStatusPlaying
}

var (
	GameStatusBetting  = GameStatus{"betting"}
	GameStatusPlaying  = GameStatus{"playing"}
	GameStatusFinished = GameStatus{"finished"}
	GameStatuses       = []GameStatus{GameStatusBetting, GameStatusPlaying, GameStatusFinished}
)

func NewGameStatus(statusString string) GameStatus {
	for _, status := range GameStatuses {
		if status.value == statusString {
			return status
		}
	}
	panic("Invalid status: " + statusString)
}
