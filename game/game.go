package game

import (
	"encoding/gob"
	"github.com/michalslomczynski/go-topdown-mmo-client/network/protocols"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/entities/player"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/entities/playerlist"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/worldmap"
	"log"
	"net"
	"sync"
)

type Game struct {
	GameMap *worldmap.Map
	Hero    *player.Player
	Players playerlist.PlayerList
	Dec     *gob.Decoder
	Enc     *gob.Encoder
	sync.RWMutex
}

func (g *Game) Init() {
	conn, err := net.Dial("tcp", protocols.ServerPort)
	if err != nil {
		log.Fatal(err)
	}
	g.Dec = gob.NewDecoder(conn)
	g.Enc = gob.NewEncoder(conn)

	worldmap.Worldmap = *protocols.AskForMap(g.Dec, g.Enc)
	g.GameMap = &worldmap.Worldmap

	player.Hero = *protocols.AskForNewPlayer(g.Dec, g.Enc)
	player.Hero.X = 15
	player.Hero.Y = 15
	player.Hero.Level = 10
	g.Hero = &player.Hero

	g.Players = playerlist.Players

	g.GameMap.Init()

	g.Hero.Name = "nickname"
}
