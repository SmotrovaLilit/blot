package gameset

import "blot/internal/blot/domain/gameset/player"

type SittingOrder [4]player.ID

func NewPlayersSittingOrder(u1, u2, u3, u4 player.ID) SittingOrder {
	return [4]player.ID{u1, u2, u3, u4}
}
