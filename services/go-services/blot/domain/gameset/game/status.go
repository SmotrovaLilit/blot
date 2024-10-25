package game

type Status struct {
	value string
}

func (s Status) IsFinished() bool {
	return s == GameStatusFinished
}

func (s Status) CanPlayCard() bool {
	return s == GameStatusPlaying
}

func (s Status) IsZero() bool {
	return s.value == ""
}

func (s Status) String() string {
	return s.value
}

var (
	GameStatusBetting  = Status{"betting"}
	GameStatusPlaying  = Status{"playing"}
	GameStatusFinished = Status{"finished"}
	GameStatuses       = []Status{GameStatusBetting, GameStatusPlaying, GameStatusFinished}
)

func NewStatus(statusString string) Status {
	for _, status := range GameStatuses {
		if status.value == statusString {
			return status
		}
	}
	panic("Invalid status: " + statusString)
}
