package playerlist

import (
	"encoding/gob"
	"fmt"
	"github.com/michalslomczynski/go-topdown-mmo-client/network/protocols"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/entities/player"
	"sync"
	"time"
)

type PlayerList map[int16]*player.Player

var Players PlayerList = make(map[int16]*player.Player)

func Update(m *sync.RWMutex, dec *gob.Decoder, enc *gob.Encoder) time.Duration {
	t := time.Now()
	list := *protocols.AskForPlayerPosList(dec, enc)
	s := time.Since(t)

	if list == nil {
		return s
	}

	m.Lock()
	for key, pos := range list {
		if _, ok := Players[key]; !ok {
			Players[key] = &player.Player{key, pos.X, pos.Y, pos.MovingDirection, 10, 0, 0, 0, fmt.Sprintf("Player%v", key)}
		} else {
			Players[key].MovingDirection = pos.MovingDirection
			Players[key].X = pos.X
			Players[key].Y = pos.Y
		}
	}
	m.Unlock()
	return s
}
