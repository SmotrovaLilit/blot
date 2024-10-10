package gameset

type SittingOrder [4]*Player

func NewPlayersSittingOrder(u1, u2, u3, u4 *Player) SittingOrder {
	return [4]*Player{u1, u2, u3, u4}
}
