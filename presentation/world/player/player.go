package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/entities/player"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/worldmap"
)

func Display(p *player.Player, screen *ebiten.Image, m *worldmap.Map, hero *player.Player) {
	mx := m.ScreenWidth / m.TileSize / 2
	my := m.ScreenHeight / m.TileSize / 2
	dx := p.X - hero.X + mx
	dy := p.Y - hero.Y + my
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(int(dx)*int(m.TileSize))+p.Ux*float64(m.TileSize)-hero.Ux*float64(m.TileSize),
		float64(int(dy)*int(m.TileSize))+p.Uy*float64(m.TileSize)-hero.Uy*float64(m.TileSize))
	screen.DrawImage(m.GetTile(int16(375)), op)
}

func DisplayShadow(p *player.Player, screen *ebiten.Image, m *worldmap.Map, hero *player.Player) {
	mx := m.ScreenWidth / m.TileSize / 2
	my := m.ScreenHeight / m.TileSize / 2
	dx := p.X - hero.X + mx
	dy := p.Y - hero.Y + my
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(int(dx)*int(m.TileSize))+p.Ux*float64(m.TileSize)-hero.Ux*float64(m.TileSize),
		float64(int(dy)*int(m.TileSize))+p.Uy*float64(m.TileSize)-hero.Uy*float64(m.TileSize))
	screen.DrawImage(m.GetTile(int16(375)), op)
}
