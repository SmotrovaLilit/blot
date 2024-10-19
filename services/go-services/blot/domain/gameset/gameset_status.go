package gameset

var (
	GamesetStatusWaitedForPlayers = GamesetStatus{"waited_for_players"}
	GamesetStatusReadyToStart     = GamesetStatus{"ready_to_start"}
	GamesetStatuses               = []GamesetStatus{GamesetStatusWaitedForPlayers, GamesetStatusReadyToStart}
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

func (s GamesetStatus) CanJoin() bool {
	return s == GamesetStatusWaitedForPlayers
}
