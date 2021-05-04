package _map

import (
	"fmt"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/entities/player"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/worldmap"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawMap(hero *player.Player, m *worldmap.Map, screen *ebiten.Image, x, y int) {
	for _, l := range m.Layers {
		for i := -1; i <= int(m.XNumInScreen); i++ {
			for j := -1; j <= int(m.YNumInScreen); j++ {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(i*int(m.TileSize))-hero.Ux*(float64(m.TileSize)), float64(j*int(m.TileSize))-hero.Uy*float64(m.TileSize))
				t := l[(x+i-int(m.ScreenWidth/m.TileSize/2))+(j+y-int(m.ScreenHeight/m.TileSize/2))*int(m.MapWidth)]
				screen.DrawImage(m.GetTile(t), op)
			}
		}
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
	}
}
