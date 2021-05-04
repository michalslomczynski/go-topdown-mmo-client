package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/entities/player"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/entities/playerlist"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/worldmap"
	"math"
	"sync"
)

var Shadows = make(map[int16]*player.Player)

func DisplayShadows(hero *player.Player, list playerlist.PlayerList, m *worldmap.Map, mu *sync.RWMutex, screen *ebiten.Image) {
	mu.RLock()
	for key, pl := range list {
		if playerInRadius(pl, hero, m) {
			if _, ok := Shadows[key]; !ok {
				Shadows[key] = &player.Player{
					pl.ID,
					pl.X,
					pl.Y,
					pl.MovingDirection,
					pl.Level,
					pl.Outfit,
					0,
					0,
					pl.Name}
			}
			UpdateShadow(Shadows[key], pl)
			if key != hero.ID {
				DisplayShadow(Shadows[key], screen, m, hero)
			}
		} else if _, ok := Shadows[key]; ok {
			delete(Shadows, key)
		}
	}
	mu.RUnlock()
}

func UpdateShadow(p *player.Player, pl *player.Player) {
	switch p.MovingDirection {

	case player.None:
		p.Ux = 0
		p.Uy = 0

	case player.Left:
		p.Ux -= 0.01 * float64(p.Level)

	case player.Right:
		p.Ux += 0.01 * float64(p.Level)

	case player.Down:
		p.Uy += 0.01 * float64(p.Level)

	case player.Up:
		p.Uy -= 0.01 * float64(p.Level)
	}
	StopShadow(p)
	SyncShadow(p, pl)
}

func playerInRadius(p *player.Player, hero *player.Player, m *worldmap.Map) bool {
	if (math.Abs(float64(p.X-hero.X)) < float64(m.XNumInScreen)) && (math.Abs(float64(p.Y-hero.Y)) < float64(m.YNumInScreen)) {
		return true
	}
	return false
}

func StopShadow(p *player.Player) {
	if p.Ux >= 1 {
		p.X++
		p.Ux = 0
		p.MovingDirection = player.None
	} else if p.Ux <= -1 {
		p.X--
		p.Ux = 0
		p.MovingDirection = player.None
	} else if p.Uy >= 1 {
		p.Y++
		p.Uy = 0
		p.MovingDirection = player.None
	} else if p.Uy <= -1 {
		p.Y--
		p.Uy = 0
		p.MovingDirection = player.None
	}
}

func SyncShadow(p *player.Player, pl *player.Player) {
	if p.MovingDirection == player.None {
		if pl.X-p.X >= 1 {
			p.MovingDirection = player.Right
		} else if pl.X-p.X <= -1 {
			p.MovingDirection = player.Left
		} else if pl.Y-p.Y >= 1 {
			p.MovingDirection = player.Down
		} else if pl.Y-p.Y <= -1 {
			p.MovingDirection = player.Up
		}
	}
}
