package self

import (
	"encoding/gob"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/michalslomczynski/go-topdown-mmo-client/network/protocols"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/entities/player"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/worldmap"
)

func Display(screen *ebiten.Image, m *worldmap.Map, p *player.Player) {
	mx := m.ScreenWidth / m.TileSize / 2
	my := m.ScreenHeight / m.TileSize / 2
	dx := mx
	dy := my
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(int(dx)*int(m.TileSize)), float64(int(dy)*int(m.TileSize)))
	screen.DrawImage(m.GetTile(int16(375)), op)
}

func Move(p *player.Player, enc *gob.Encoder) {
	switch p.MovingDirection {

	case player.Left:
		p.Ux -= 0.01 * float64(p.Level)
		if p.Ux <= -1 {
			Stop(p, enc)
			p.Ux = 0
			p.X--
		}

	case player.Right:
		p.Ux += 0.01 * float64(p.Level)
		if p.Ux >= 1 {
			Stop(p, enc)
			p.Ux = 0
			p.X++
		}

	case player.Up:
		p.Uy -= 0.01 * float64(p.Level)
		if p.Uy <= -1 {
			Stop(p, enc)
			p.Uy = 0
			p.Y--
		}

	case player.Down:
		p.Uy += 0.01 * float64(p.Level)
		if p.Uy >= 1 {
			Stop(p, enc)
			p.Uy = 0
			p.Y++
		}
	}
}

func Stop(p *player.Player, enc *gob.Encoder) {
	p.MovingDirection = player.None
	protocols.AskForStopPlayer(enc, p.ID)
}
