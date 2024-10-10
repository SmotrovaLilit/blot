package gameset

type ErrGameSetCannotCreateGame struct {
	GameSetID ID
}

func (e ErrGameSetCannotCreateGame) Error() string {
	return "game set " + e.GameSetID.String() + " cannot create game"
}

var (
	StatusWaitedToStartGame = Status{"waited_to_start_game"}
	StatusPlaying           = Status{"playing"}
	StatusFinished          = Status{"finished"}
	StatusCanceled          = Status{"canceled"}
	Statuses                = []Status{StatusWaitedToStartGame, StatusPlaying, StatusFinished, StatusCanceled}
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

func (s Status) CanStartNewGame() bool {
	return s == StatusWaitedToStartGame
}
