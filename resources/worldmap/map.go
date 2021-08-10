package worldmap

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/michalslomczynski/go-topdown-mmo-client/data/window"

	"github.com/hajimehoshi/ebiten/v2"
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

var (
	Worldmap Map
	//go:embed tiles.png
	tilesBytes []byte
)

func (m *Map) Init() {
	m.ScreenHeight = window.ScreenHeight
	m.ScreenWidth = window.ScreenWidth
	m.XNumInScreen = m.ScreenWidth / m.TileSize
	m.YNumInScreen = m.ScreenHeight / m.TileSize

	tiles, _, err := image.Decode(bytes.NewReader(tilesBytes))
	if err != nil {
		log.Fatal(err)
	}
	m.tilesImage = ebiten.NewImageFromImage(tiles)
}

func (m *Map) GetTile(val int16) *ebiten.Image {
	y := val / m.TileXNum * m.TileSize
	x := val % m.TileXNum * m.TileSize
	return m.tilesImage.SubImage(image.Rect(int(x), int(y), int(x+m.TileSize), int(y+m.TileSize))).(*ebiten.Image)
}
