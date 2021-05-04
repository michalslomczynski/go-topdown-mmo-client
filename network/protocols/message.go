package protocols

const (
	ServerPort = ":8081"
)

const (
	OverAndOut = iota
	MapRequest
	NewPlayerRequest
	PlayerListRequest
	PlayerPosListRequest
	MovePlayerLeftRequest
	MovePlayerRightRequest
	MovePlayerUpRequest
	MovePlayerDownRequest
	StopPlayerRequest
)

type Message struct {
	ID   int
	Data []byte
}
