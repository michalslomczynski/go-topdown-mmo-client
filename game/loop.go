package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/michalslomczynski/go-topdown-mmo-client/data/window"
	"github.com/michalslomczynski/go-topdown-mmo-client/input"
	"github.com/michalslomczynski/go-topdown-mmo-client/presentation/world/player"
	"github.com/michalslomczynski/go-topdown-mmo-client/presentation/world/self"
	"github.com/michalslomczynski/go-topdown-mmo-client/presentation/world/view"
)

func (g *Game) reset() {
}

func (g *Game) Update() error {
	input.HandleInput(g.Hero, g.GameMap, g.Enc)
	self.Move(g.Hero, g.Enc)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	view.Display(screen, g.GameMap, g.Hero, g.Players)
	player.DisplayShadows(g.Hero, g.Players, g.GameMap, &g.RWMutex, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return window.ScreenWidth, window.ScreenHeight
}
