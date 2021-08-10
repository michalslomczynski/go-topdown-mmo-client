package mobile

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"
	"github.com/michalslomczynski/go-topdown-mmo-client/game"
)

func init() {
	mobile.SetGame(&game.Game{})
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
