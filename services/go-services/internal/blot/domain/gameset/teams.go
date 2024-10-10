package gameset

//
//import (
//	"errors"
//	"blot/internal/blot/domain/user"
//
//	"blot/internal/blot/domain/team"
//)
//
//var ErrSameTeam = errors.New("same team")
//var ErrSamePlayer = errors.New("same player")
//
//type Teams struct {
//	values [2]Team
//}
//
//func (t Teams) First() Team {
//	return t.values[0]
//}
//
//func (t Teams) Second() Team {
//	return t.values[1]
//}
//
//func NewTeams(f, s Team) (Teams, error) {
//	if f.IsZero() || s.IsZero() {
//		panic("empty team, use constructor to create object")
//	}
//	if f.Equal(s) {
//		return Teams{}, ErrSameTeam
//	}
//	return Teams{values: [2]Team{f, s}}, nil
//}
//
//type Team struct {
//	id      team.ID
//	players [2]user.ID
//}
//
//func (t Team) Equal(s Team) bool {
//	return t.id == s.id
//}
//
//func (t Team) IsZero() bool {
//	return t == Team{}
//}
//
//func (t Team) ID() team.ID {
//	return t.id
//}
//
//func (t Team) Players() [2]user.ID {
//	return t.players
//}
//
//func (t Team) FirstPlayer() user.ID {
//	return t.players[0]
//}
//
//func (t Team) SecondPlayer() user.ID {
//	return t.players[1]
//}
//
//func NewTeam(id team.ID, p1, p2 user.ID) Team {
//	if id.IsZero() || p1.IsZero() || p2.IsZero() {
//		panic("empty objects, use constructor to create object")
//	}
//	if p1.Equal(p2) {
//		panic("same player")
//	}
//	return Team{id: id, players: [2]user.ID{p1, p2}}
//}
