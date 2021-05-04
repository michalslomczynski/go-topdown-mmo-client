package worldmap

import (
	"github.com/michalslomczynski/go-topdown-mmo-client/data/window"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Map struct {
	TileSize  int16
	TileXNum  int16
	MapHeight int16
	MapWidth  int16
	Layers    [][]int16

	/* client resources below */
	XNumInScreen int16
	YNumInScreen int16
	ScreenWidth  int16
	ScreenHeight int16
	tilesImage   *ebiten.Image
}

type Position struct {
	X int16
	Y int16
}

var Worldmap Map

func (m *Map) Init() {
	m.ScreenHeight = window.ScreenHeight
	m.ScreenWidth = window.ScreenWidth
	m.XNumInScreen = m.ScreenWidth / m.TileSize
	m.YNumInScreen = m.ScreenHeight / m.TileSize

	var err error
	m.tilesImage, _, err = ebitenutil.NewImageFromFile("resources/assets/map/tiles.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Map) GetTile(val int16) *ebiten.Image {
	y := val / m.TileXNum * m.TileSize
	x := val % m.TileXNum * m.TileSize
	return m.tilesImage.SubImage(image.Rect(int(x), int(y), int(x+m.TileSize), int(y+m.TileSize))).(*ebiten.Image)
}
