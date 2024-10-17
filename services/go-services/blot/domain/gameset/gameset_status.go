package gameset

var (
	GamesetStatusWaitedForPlayers = GamesetStatus{"waited_for_players"}
	GamesetStatuses               = []GamesetStatus{GamesetStatusWaitedForPlayers}
)

type GamesetStatus struct {
	value string
}

func NewGamesetStatus(value string) GamesetStatus {
	for _, s := range GamesetStatuses {
		if s.value == value {
			return s
		}
	}
	panic("invalid status")
}

func (s GamesetStatus) String() string {
	return s.value
}
