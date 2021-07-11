package protocols

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/michalslomczynski/go-topdown-mmo-client/resources/entities/player"
	"github.com/michalslomczynski/go-topdown-mmo-client/resources/worldmap"
)

func AskForMap(dec *gob.Decoder, enc *gob.Encoder) *worldmap.Map {
	msg := Message{MapRequest, make([]byte, 0)}
	enc.Encode(msg)

	m := new(worldmap.Map)
	dec.Decode(m)

	return m
}

func AskForNewPlayer(dec *gob.Decoder, enc *gob.Encoder) *player.Player {
	msg := Message{NewPlayerRequest, make([]byte, 0)}
	enc.Encode(msg)

	p := new(player.Player)
	dec.Decode(p)

	return p
}

func AskForPlayerList(dec *gob.Decoder, enc *gob.Encoder) *map[int16]player.Player {
	msg := Message{PlayerListRequest, make([]byte, 0)}
	enc.Encode(msg)

	pl := new(map[int16]player.Player)
	dec.Decode(pl)

	return pl
}

func AskForPlayerPosList(dec *gob.Decoder, enc *gob.Encoder) *map[int16]player.Position {
	msg := Message{PlayerPosListRequest, make([]byte, 0)}
	enc.Encode(msg)

	pl := new(map[int16]player.Position)
	dec.Decode(pl)

	return pl
}

func AskForMovePlayerLeft(enc *gob.Encoder, ID int16) {
	buff := new(bytes.Buffer)
	objenc := gob.NewEncoder(buff)
	err := objenc.Encode(ID)

	if err != nil {
		log.Fatal(err)
	}

	msg := Message{MovePlayerLeftRequest, buff.Bytes()}
	enc.Encode(msg)
}

func AskForMovePlayerRight(enc *gob.Encoder, ID int16) {
	buff := new(bytes.Buffer)
	objenc := gob.NewEncoder(buff)
	err := objenc.Encode(ID)

	if err != nil {
		log.Fatal(err)
	}

	msg := Message{MovePlayerRightRequest, buff.Bytes()}
	enc.Encode(msg)
}

func AskForMovePlayerUp(enc *gob.Encoder, ID int16) {
	buff := new(bytes.Buffer)
	objenc := gob.NewEncoder(buff)
	err := objenc.Encode(ID)

	if err != nil {
		log.Fatal(err)
	}

	msg := Message{MovePlayerUpRequest, buff.Bytes()}
	enc.Encode(msg)
}

func AskForMovePlayerDown(enc *gob.Encoder, ID int16) {
	buff := new(bytes.Buffer)
	objenc := gob.NewEncoder(buff)
	err := objenc.Encode(ID)

	if err != nil {
		log.Fatal(err)
	}

	msg := Message{MovePlayerDownRequest, buff.Bytes()}
	enc.Encode(msg)
}

func AskForStopPlayer(enc *gob.Encoder, ID int16) {
	buff := new(bytes.Buffer)
	objenc := gob.NewEncoder(buff)
	err := objenc.Encode(ID)

	if err != nil {
		log.Fatal(err)
	}

	msg := Message{StopPlayerRequest, buff.Bytes()}
	enc.Encode(msg)
}
