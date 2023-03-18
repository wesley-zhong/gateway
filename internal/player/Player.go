package player

import (
	"gateway/internal/client"
)

//------player

type Player struct {
	Context *client.ConnClientContext
	Pid     int64
}

func NewPlayer(pid int64, context *client.ConnClientContext) *Player {
	return &Player{Context: context, Pid: pid}
}

// player mgr----

func NewPlayerMgr() *PlayerMgr {
	return &PlayerMgr{
		playerIdMap:  make(map[int64]*Player),
		playerSidMap: make(map[int64]*Player),
	}
}

type PlayerMgr struct {
	playerIdMap  map[int64]*Player
	playerSidMap map[int64]*Player
}

func (playerMgr *PlayerMgr) Add(player *Player) {
	playerMgr.playerIdMap[int64(player.Pid)] = player
	playerMgr.playerSidMap[player.Context.Sid] = player
}

func (playerMgr *PlayerMgr) GetByRoleId(pid int64) *Player {
	return playerMgr.playerIdMap[pid]
}

func (playerMgr *PlayerMgr) GetByContext(context *client.ConnClientContext) *Player {
	return playerMgr.playerIdMap[context.Sid]
}
