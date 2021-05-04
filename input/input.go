package input

import (
	"encoding/gob"
	"github.com/michalslomczynski/go-topdown-mmo-client/network/protocols"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/entities/player"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/worldmap"

	"github.com/hajimehoshi/ebiten/v2"
)

func HandleInput(p *player.Player, m *worldmap.Map, enc *gob.Encoder) {
	if p.MovingDirection == player.None {
		if (ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA)) && (getCollisionTile(m, (*p).X-1, (*p).Y) <= 0) && ((*p).X > 0) {
			protocols.AskForMovePlayerLeft(enc, (*p).ID)
			p.MovingDirection = player.Left
		} else if (ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD)) && (getCollisionTile(m, (*p).X+1, (*p).Y) <= 0) && ((*p).X < m.MapWidth) {
			protocols.AskForMovePlayerRight(enc, (*p).ID)
			p.MovingDirection = player.Right
		} else if (ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS)) && (getCollisionTile(m, (*p).X, (*p).Y+1) <= 0) && ((*p).Y < m.MapHeight) {
			protocols.AskForMovePlayerDown(enc, (*p).ID)
			p.MovingDirection = player.Down
		} else if (ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW)) && (getCollisionTile(m, (*p).X, (*p).Y-1) <= 0) && ((*p).Y > 0) {
			protocols.AskForMovePlayerUp(enc, (*p).ID)
			p.MovingDirection = player.Up
		}
	}
}

func getCollisionTile(m *worldmap.Map, x, y int16) int16 {
	return m.Layers[1][int(x)+int(y)*int(m.MapWidth)]
}
