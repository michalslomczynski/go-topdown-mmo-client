package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/michalslomczynski/go-topdown-mmo-client/data/window"
	"github.com/michalslomczynski/go-topdown-mmo-client/game"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/entities/playerlist"
	"log"
)

func newGame() *game.Game {
	g := &game.Game{}
	g.Init()
	return g
}

func main() {
	ebiten.SetWindowSize(window.ScreenWidth*3, window.ScreenHeight*3)
	ebiten.SetWindowTitle("MMORPG Demo")
	newGame := newGame()

	go func() {
		for {
			playerlist.Update(&newGame.RWMutex, newGame.Dec, newGame.Enc)
		}
	}()

	if err := ebiten.RunGame(newGame); err != nil {
		log.Fatal(err)
	}
}
