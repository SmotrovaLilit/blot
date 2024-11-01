package game

type Status struct {
	value string
}

var (
	StatusBetting  = Status{"betting"}
	StatusPlaying  = Status{"playing"}
	StatusFinished = Status{"finished"}
	Statuses       = []Status{StatusBetting, StatusPlaying, StatusFinished}
)

func NewStatus(statusString string) Status {
	for _, status := range Statuses {
		if status.value == statusString {
			return status
		}
	}
	panic("Invalid status: " + statusString)
}

func (s Status) IsFinished() bool {
	return s == StatusFinished
}

func (s Status) CanPlayCard() bool {
	return s == StatusPlaying
}

func (s Status) IsZero() bool {
	return s.value == ""
}

func (s Status) String() string {
	return s.value
}

func (s Status) SetBet() (Status, error) {
	if s != StatusBetting {
		return Status{}, ErrGameNotReadyToSetBet{
			Status: s.String(),
		}
	}

	return StatusPlaying, nil
}

type ErrGameNotReadyToSetBet struct {
	Status string
}

func (e ErrGameNotReadyToSetBet) Error() string {
	return "game is not ready to set bet, current status: " + e.Status
}
