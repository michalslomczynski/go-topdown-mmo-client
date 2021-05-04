package view

import (
	"github.com/hajimehoshi/ebiten/v2"
	_map "github.com/michalslomczynski/go-topdown-mmo-client/presentation/world/map"
	"github.com/michalslomczynski/go-topdown-mmo-client/presentation/world/self"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/entities/player"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/entities/playerlist"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/worldmap"
	"math"
)

func Display(screen *ebiten.Image, m *worldmap.Map, hero *player.Player, players playerlist.PlayerList) {
	_map.DrawMap(hero, m, screen, int(hero.X), int(hero.Y))
	self.Display(screen, m, hero)
}

func playerInRadius(p *player.Player, hero *player.Player, m *worldmap.Map) bool {
	if (math.Abs(float64(p.X-hero.X)) < float64(m.XNumInScreen)) && (math.Abs(float64(p.Y-hero.Y)) < float64(m.YNumInScreen)) {
		return true
	}
	return false
}
